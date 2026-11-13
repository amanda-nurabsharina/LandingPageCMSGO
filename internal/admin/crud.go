package admin

import (
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/mail"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"be-lp-go/internal/models"
	"be-lp-go/internal/seed"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// FormState carries values & validation errors to the form template.
type FormState struct {
	Values map[string]string
	Tags   map[string][]string
	Images map[string]string
	Errors map[string]string
}

func newFormState() FormState {
	return FormState{
		Values: map[string]string{},
		Tags:   map[string][]string{},
		Images: map[string]string{},
		Errors: map[string]string{},
	}
}

type header struct {
	Label    string
	URL      string
	Sortable bool
	Active   bool
	Dir      string
}

type cell struct {
	Kind ColKind
	Text string
	URL  string
}

type row struct {
	ID      string
	EditURL string
	Cells   []cell
}

type card struct {
	Label string
	Icon  string
	URL   string
	Count int64
}

// ---------------------------------------------------------------------------
// Dashboard
// ---------------------------------------------------------------------------

func (a *Admin) dashboard(c *fiber.Ctx) error {
	var cards []card
	for _, r := range Resources {
		if r.Singleton {
			continue
		}
		var n int64
		a.db.Table(r.Table).Count(&n)
		cards = append(cards, card{Label: r.Label, Icon: r.Icon, URL: "/admin/" + r.Slug, Count: n})
	}
	var settings []card
	for _, r := range Resources {
		if r.Singleton {
			settings = append(settings, card{Label: r.Label, Icon: r.Icon, URL: "/admin/" + r.Slug})
		}
	}
	return a.render(c, "dashboard", fiber.Map{
		"Title":    "Dashboard",
		"Active":   "dashboard",
		"Cards":    cards,
		"Settings": settings,
	})
}

// ---------------------------------------------------------------------------
// List
// ---------------------------------------------------------------------------

func (a *Admin) list(c *fiber.Ctx) error {
	r := findResource(c.Params("slug"))
	if r == nil {
		return fiber.ErrNotFound
	}
	if r.Singleton {
		if _, _, _, _, _, _, err := seed.Singletons(a.db); err != nil {
			return err
		}
		return c.Redirect(fmt.Sprintf("/admin/%s/1/edit", r.Slug))
	}

	q := strings.TrimSpace(c.Query("q"))
	sortCol := c.Query("sort")
	dir := strings.ToLower(c.Query("dir"))
	if dir != "desc" {
		dir = "asc"
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))
	if page < 1 {
		page = 1
	}

	scope := func(tx *gorm.DB) *gorm.DB {
		if q == "" {
			return tx
		}
		var conds []string
		var args []any
		for _, col := range r.Columns {
			if col.Searchable {
				conds = append(conds, fmt.Sprintf(`CAST("%s" AS TEXT) ILIKE ?`, col.Name))
				args = append(args, "%"+q+"%")
			}
		}
		if len(conds) == 0 {
			return tx
		}
		return tx.Where(strings.Join(conds, " OR "), args...)
	}

	var total int64
	if err := scope(a.db.Table(r.Table)).Count(&total).Error; err != nil {
		return err
	}

	order := "id ASC"
	if col := r.column(sortCol); col != nil && col.Sortable {
		order = fmt.Sprintf(`"%s" %s, id ASC`, col.Name, strings.ToUpper(dir))
	} else {
		sortCol = ""
		if r.DefaultSort != "" {
			order = fmt.Sprintf(`"%s" ASC, id ASC`, r.DefaultSort)
		}
	}

	pages := int((total + perPage - 1) / perPage)
	if pages < 1 {
		pages = 1
	}
	if page > pages {
		page = pages
	}

	var records []map[string]any
	if err := scope(a.db.Table(r.Table)).Order(order).Limit(perPage).Offset((page - 1) * perPage).Find(&records).Error; err != nil {
		return err
	}

	base := "/admin/" + r.Slug
	buildURL := func(sort, d string, p int) string {
		v := url.Values{}
		if q != "" {
			v.Set("q", q)
		}
		if sort != "" {
			v.Set("sort", sort)
			v.Set("dir", d)
		}
		if p > 1 {
			v.Set("page", strconv.Itoa(p))
		}
		if len(v) == 0 {
			return base
		}
		return base + "?" + v.Encode()
	}

	var headers []header
	for _, col := range r.Columns {
		h := header{Label: col.DisplayLabel(), Sortable: col.Sortable}
		if col.Sortable {
			next := "asc"
			if sortCol == col.Name {
				h.Active = true
				h.Dir = dir
				if dir == "asc" {
					next = "desc"
				}
			}
			h.URL = buildURL(col.Name, next, 1)
		}
		headers = append(headers, h)
	}

	var rows []row
	for _, rec := range records {
		id := toString(rec["id"])
		rw := row{ID: id, EditURL: fmt.Sprintf("%s/%s/edit", base, id)}
		for _, col := range r.Columns {
			v := toString(rec[col.Name])
			cl := cell{Kind: col.Kind, Text: v}
			switch col.Kind {
			case CImage:
				if v != "" {
					cl.URL = "/storage/" + v
				}
			case CBool:
				if v == "true" || v == "1" || v == "t" {
					cl.Text = "Aktif"
				} else {
					cl.Text = "Nonaktif"
				}
			case CStars:
				n, _ := strconv.Atoi(v)
				if n < 0 {
					n = 0
				}
				if n > 5 {
					n = 5
				}
				cl.Text = strings.Repeat("â˜…", n) + strings.Repeat("â˜†", 5-n)
			}
			rw.Cells = append(rw.Cells, cl)
		}
		rows = append(rows, rw)
	}

	data := fiber.Map{
		"Title":    r.Label,
		"Active":   r.Slug,
		"Resource": r,
		"Headers":  headers,
		"Rows":     rows,
		"Query":    q,
		"Total":    total,
		"Page":     page,
		"Pages":    pages,
	}
	if page > 1 {
		data["PrevURL"] = buildURL(sortCol, dir, page-1)
	}
	if page < pages {
		data["NextURL"] = buildURL(sortCol, dir, page+1)
	}
	return a.render(c, "list", data)
}

// ---------------------------------------------------------------------------
// Create / Store
// ---------------------------------------------------------------------------

func (a *Admin) createPage(c *fiber.Ctx) error {
	r := findResource(c.Params("slug"))
	if r == nil {
		return fiber.ErrNotFound
	}
	if r.Singleton {
		return c.Redirect("/admin/" + r.Slug)
	}
	st := newFormState()
	for _, f := range r.Fields {
		switch f.Type {
		case FTags:
			st.Tags[f.Name] = append([]string{}, f.DefaultTags...)
		case FToggle:
			if f.Default != "" {
				st.Values[f.Name] = f.Default
			} else {
				st.Values[f.Name] = "true"
			}
		case FImage:
		default:
			st.Values[f.Name] = f.Default
		}
	}
	return a.renderForm(c, r, st, 0)
}

func (a *Admin) store_(c *fiber.Ctx) error {
	r := findResource(c.Params("slug"))
	if r == nil {
		return fiber.ErrNotFound
	}
	if r.Singleton {
		return c.Redirect("/admin/" + r.Slug)
	}

	data, st, newFiles, _, ok, err := a.parseForm(c, r, map[string]any{})
	if err != nil {
		return err
	}
	if !ok {
		return a.renderForm(c.Status(fiber.StatusUnprocessableEntity), r, st, 0)
	}

	now := time.Now()
	data["created_at"] = now
	data["updated_at"] = now
	if err := a.db.Table(r.Table).Create(data).Error; err != nil {
		for _, f := range newFiles {
			a.storage.Delete(f)
		}
		return err
	}
	a.setFlash(c, r.Singular+" berhasil dibuat.")
	return c.Redirect("/admin/" + r.Slug)
}

// ---------------------------------------------------------------------------
// Edit / Update
// ---------------------------------------------------------------------------

func (a *Admin) loadRecord(c *fiber.Ctx) (*Resource, int, map[string]any, error) {
	r := findResource(c.Params("slug"))
	if r == nil {
		return nil, 0, nil, fiber.ErrNotFound
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return nil, 0, nil, fiber.ErrNotFound
	}
	if r.Singleton && id == 1 {
		if _, _, _, _, _, _, err := seed.Singletons(a.db); err != nil {
			return nil, 0, nil, err
		}
	}
	rec := map[string]any{}
	if err := a.db.Table(r.Table).Where("id = ?", id).Take(&rec).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, nil, fiber.ErrNotFound
		}
		return nil, 0, nil, err
	}
	if len(rec) == 0 {
		return nil, 0, nil, fiber.ErrNotFound
	}
	return r, id, rec, nil
}

func (a *Admin) editPage(c *fiber.Ctx) error {
	r, id, rec, err := a.loadRecord(c)
	if err != nil {
		return err
	}
	st := newFormState()
	for _, f := range r.Fields {
		switch f.Type {
		case FTags:
			st.Tags[f.Name] = parseTags(rec[f.Name])
		case FImage:
			st.Images[f.Name] = toString(rec[f.Name])
		default:
			st.Values[f.Name] = toString(rec[f.Name])
		}
	}
	return a.renderForm(c, r, st, id)
}

func (a *Admin) update(c *fiber.Ctx) error {
	r, id, rec, err := a.loadRecord(c)
	if err != nil {
		return err
	}

	data, st, newFiles, oldFiles, ok, err := a.parseForm(c, r, rec)
	if err != nil {
		return err
	}
	if !ok {
		return a.renderForm(c.Status(fiber.StatusUnprocessableEntity), r, st, id)
	}

	data["updated_at"] = time.Now()
	if err := a.db.Table(r.Table).Where("id = ?", id).Updates(data).Error; err != nil {
		for _, f := range newFiles {
			a.storage.Delete(f)
		}
		return err
	}
	for _, f := range oldFiles {
		a.storage.Delete(f)
	}
	a.setFlash(c, r.Singular+" berhasil disimpan.")
	return c.Redirect(fmt.Sprintf("/admin/%s/%d/edit", r.Slug, id))
}

// ---------------------------------------------------------------------------
// Delete
// ---------------------------------------------------------------------------

func (a *Admin) destroy(c *fiber.Ctx) error {
	r, id, rec, err := a.loadRecord(c)
	if err != nil {
		return err
	}
	if r.Singleton {
		return fiber.NewError(fiber.StatusForbidden, "Data pengaturan tidak dapat dihapus.")
	}
	if err := a.db.Exec(fmt.Sprintf(`DELETE FROM "%s" WHERE id = ?`, r.Table), id).Error; err != nil {
		return err
	}
	for _, f := range r.imageFields() {
		a.storage.Delete(toString(rec[f.Name]))
	}
	a.setFlash(c, r.Singular+" berhasil dihapus.")
	return c.Redirect("/admin/" + r.Slug)
}

func (a *Admin) bulkDelete(c *fiber.Ctx) error {
	r := findResource(c.Params("slug"))
	if r == nil {
		return fiber.ErrNotFound
	}
	if r.Singleton {
		return fiber.NewError(fiber.StatusForbidden, "Data pengaturan tidak dapat dihapus.")
	}

	var ids []int
	for _, raw := range c.Request().PostArgs().PeekMulti("ids") {
		if n, err := strconv.Atoi(string(raw)); err == nil && n > 0 {
			ids = append(ids, n)
		}
	}
	if len(ids) == 0 {
		a.setFlash(c, "Tidak ada data yang dipilih.")
		return c.Redirect("/admin/" + r.Slug)
	}

	var records []map[string]any
	if err := a.db.Table(r.Table).Where("id IN ?", ids).Find(&records).Error; err != nil {
		return err
	}
	if err := a.db.Exec(fmt.Sprintf(`DELETE FROM "%s" WHERE id IN ?`, r.Table), ids).Error; err != nil {
		return err
	}
	for _, rec := range records {
		for _, f := range r.imageFields() {
			a.storage.Delete(toString(rec[f.Name]))
		}
	}
	a.setFlash(c, fmt.Sprintf("%d data berhasil dihapus.", len(records)))
	return c.Redirect("/admin/" + r.Slug)
}

// ---------------------------------------------------------------------------
// Form rendering & parsing
// ---------------------------------------------------------------------------

func (a *Admin) renderForm(c *fiber.Ctx, r *Resource, st FormState, id int) error {
	data := fiber.Map{
		"Active":   r.Slug,
		"Resource": r,
		"Form":     st,
		"IsEdit":   id > 0,
	}
	if id > 0 {
		data["Title"] = "Edit " + r.Singular
		data["Action"] = fmt.Sprintf("/admin/%s/%d", r.Slug, id)
		if !r.Singleton {
			data["DeleteURL"] = fmt.Sprintf("/admin/%s/%d/delete", r.Slug, id)
		}
	} else {
		data["Title"] = "Create " + r.Singular
		data["Action"] = "/admin/" + r.Slug
	}
	if len(st.Errors) > 0 {
		data["Flash"] = ""
		data["FormError"] = "Periksa kembali isian yang ditandai."
	}
	return a.render(c, "form", data)
}

// parseForm validates the submitted form and returns the column => value map.
// newFiles are freshly stored uploads, oldFiles are files to remove after a
// successful save.
func (a *Admin) parseForm(c *fiber.Ctx, r *Resource, existing map[string]any) (data map[string]any, st FormState, newFiles, oldFiles []string, ok bool, err error) {
	st = newFormState()
	data = map[string]any{}

	form, ferr := c.MultipartForm()
	if ferr != nil {
		form = &multipart.Form{Value: map[string][]string{}, File: map[string][]*multipart.FileHeader{}}
		c.Request().PostArgs().VisitAll(func(k, v []byte) {
			form.Value[string(k)] = append(form.Value[string(k)], string(v))
		})
	}
	get := func(name string) string {
		if v := form.Value[name]; len(v) > 0 {
			return strings.TrimSpace(v[0])
		}
		return ""
	}

	pendingUploads := map[string]*multipart.FileHeader{}

	for _, f := range r.Fields {
		label := f.DisplayLabel()
		switch f.Type {
		case FReadOnly:
			st.Values[f.Name] = toString(existing[f.Name])

		case FToggle:
			val := get(f.Name)
			isActive := val == "1" || strings.EqualFold(val, "true") || strings.EqualFold(val, "on")
			if isActive {
				st.Values[f.Name] = "true"
			} else {
				st.Values[f.Name] = "false"
			}
			data[f.Name] = isActive

		case FTags:
			var tags []string
			for _, t := range form.Value[f.Name] {
				if t = strings.TrimSpace(t); t != "" {
					tags = append(tags, t)
				}
			}
			st.Tags[f.Name] = tags
			if f.Required && len(tags) == 0 {
				st.Errors[f.Name] = label + " wajib diisi."
				continue
			}
			if tags == nil {
				tags = []string{}
			}
			data[f.Name] = models.StringList(tags)

		case FImage:
			current := toString(existing[f.Name])
			st.Images[f.Name] = current
			var fh *multipart.FileHeader
			if files := form.File[f.Name]; len(files) > 0 && files[0].Size > 0 {
				fh = files[0]
			}
			remove := get("remove_"+f.Name) == "1"
			switch {
			case fh != nil:
				pendingUploads[f.Name] = fh
			case remove && !f.Required && current != "":
				data[f.Name] = nil
				oldFiles = append(oldFiles, current)
				st.Images[f.Name] = ""
			case f.Required && current == "":
				st.Errors[f.Name] = label + " wajib diunggah."
			}

		case FNumber, FSelect:
			v := get(f.Name)
			st.Values[f.Name] = v
			if v == "" {
				if f.Required {
					st.Errors[f.Name] = label + " wajib diisi."
				} else {
					data[f.Name] = nil
				}
				continue
			}
			n, convErr := strconv.Atoi(v)
			if convErr != nil {
				st.Errors[f.Name] = label + " harus berupa angka."
				continue
			}
			if f.Type == FSelect && !hasOption(f.Options, v) {
				st.Errors[f.Name] = "Pilihan " + label + " tidak valid."
				continue
			}
			data[f.Name] = n

		default: // text, textarea, email, url
			v := get(f.Name)
			if f.Type == FTextarea {
				if raw := form.Value[f.Name]; len(raw) > 0 {
					v = strings.TrimSpace(strings.ReplaceAll(raw[0], "\r\n", "\n"))
				}
			}
			st.Values[f.Name] = v
			if v == "" {
				if f.Required {
					st.Errors[f.Name] = label + " wajib diisi."
				} else {
					data[f.Name] = nil
				}
				continue
			}
			if f.Type != FTextarea && utf8.RuneCountInString(v) > 255 {
				st.Errors[f.Name] = label + " maksimal 255 karakter."
				continue
			}
			if f.Type == FEmail {
				if addr, e := mail.ParseAddress(v); e != nil || addr.Address != v {
					st.Errors[f.Name] = label + " harus berupa alamat email yang valid."
					continue
				}
			}
			if f.Type == FURL {
				u, e := url.ParseRequestURI(v)
				if e != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
					st.Errors[f.Name] = label + " harus berupa URL yang valid (http/https)."
					continue
				}
			}
			data[f.Name] = v
		}
	}

	if len(st.Errors) > 0 {
		return data, st, nil, nil, false, nil
	}

	// Store uploads only after validation passed.
	for name, fh := range pendingUploads {
		fld := r.field(name)
		path, upErr := a.storage.SaveImage(fh, fld.Dir)
		if upErr != nil {
			st.Errors[name] = upErr.Error()
			continue
		}
		newFiles = append(newFiles, path)
		data[name] = path
		if cur := toString(existing[name]); cur != "" {
			oldFiles = append(oldFiles, cur)
		}
	}
	if len(st.Errors) > 0 {
		for _, p := range newFiles {
			a.storage.Delete(p)
		}
		return data, st, nil, nil, false, nil
	}

	return data, st, newFiles, oldFiles, true, nil
}

func hasOption(opts []Option, v string) bool {
	for _, o := range opts {
		if o.Value == v {
			return true
		}
	}
	return false
}

func toString(v any) string {
	switch t := v.(type) {
	case nil:
		return ""
	case string:
		return t
	case []byte:
		return string(t)
	case time.Time:
		return t.Format("2006-01-02 15:04")
	case *time.Time:
		if t == nil {
			return ""
		}
		return t.Format("2006-01-02 15:04")
	default:
		return fmt.Sprint(t)
	}
}

func parseTags(v any) []string {
	switch t := v.(type) {
	case []any:
		out := make([]string, 0, len(t))
		for _, x := range t {
			out = append(out, fmt.Sprint(x))
		}
		return out
	case []string:
		return t
	}
	var out []string
	if s := toString(v); s != "" {
		_ = json.Unmarshal([]byte(s), &out)
	}
	return out
}
