package admin

import "strings"

type FieldType string

const (
	FText     FieldType = "text"
	FTextarea FieldType = "textarea"
	FNumber   FieldType = "number"
	FEmail    FieldType = "email"
	FURL      FieldType = "url"
	FImage    FieldType = "image"
	FTags     FieldType = "tags"
	FSelect   FieldType = "select"
	FToggle   FieldType = "toggle"
	FReadOnly FieldType = "readonly"
)

type Option struct {
	Value string
	Label string
}

type Field struct {
	Name        string
	Label       string
	Type        FieldType
	Required    bool
	Placeholder string
	Default     string
	DefaultTags []string
	Options     []Option
	Dir         string
	Full        bool
}

func (f Field) DisplayLabel() string {
	if f.Label != "" {
		return f.Label
	}
	return humanize(f.Name)
}

type ColKind string

const (
	CText   ColKind = "text"
	CImage  ColKind = "image"
	CStars  ColKind = "stars"
	CBool   ColKind = "bool"
)

type Column struct {
	Name       string
	Label      string
	Kind       ColKind
	Sortable   bool
	Searchable bool
}

func (c Column) DisplayLabel() string {
	if c.Label != "" {
		return c.Label
	}
	return humanize(c.Name)
}

type Resource struct {
	Group       string
	Slug        string
	Label       string
	Singular    string
	Icon        string
	Table       string
	Singleton   bool
	ReadOnly    bool
	Fields      []Field
	Columns     []Column
	DefaultSort string
}

func (r *Resource) field(name string) *Field {
	for i := range r.Fields {
		if r.Fields[i].Name == name {
			return &r.Fields[i]
		}
	}
	return nil
}

func (r *Resource) column(name string) *Column {
	for i := range r.Columns {
		if r.Columns[i].Name == name {
			return &r.Columns[i]
		}
	}
	return nil
}

func (r *Resource) imageFields() []Field {
	var out []Field
	for _, f := range r.Fields {
		if f.Type == FImage {
			out = append(out, f)
		}
	}
	return out
}

func humanize(name string) string {
	s := strings.ReplaceAll(name, "_", " ")
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func sortOrderField() Field {
	return Field{Name: "sort_order", Label: "Urutan", Type: FNumber, Required: true, Default: "0"}
}

var btnURLOptions = []Option{
	{"whatsapp", "Hubungi WhatsApp"},
	{"#about", "Bagian Tentang Kami (About)"},
	{"#services", "Bagian Layanan (Services)"},
	{"#services_premium", "Bagian Layanan Premium"},
	{"#benefits", "Bagian Keunggulan (Why Choose Us)"},
	{"#portfolio", "Bagian Portofolio"},
	{"#timeline", "Bagian Cara Pesan (Order Timeline)"},
	{"#testimonials", "Bagian Testimoni"},
	{"#news", "Bagian Berita & Informasi (News)"},
	{"#activities", "Bagian Kegiatan Kami (Activities)"},
	{"#contact", "Bagian Formulir Kontak"},
}

// Resources is the full panel definition for branch General.
var Resources = []*Resource{
	// ----------------- HERO -----------------
	{
		Group: "Hero", Slug: "hero-sections", Label: "Hero Banner", Singular: "Hero Banner", Icon: "house",
		Table: "hero_sections", Singleton: true,
		Fields: []Field{
			{Name: "badge", Label: "Lencana / Teks Kecil", Type: FText, Placeholder: "contoh: Solusi Terpercaya"},
			{Name: "title", Label: "Judul Utama", Type: FText, Required: true, Default: "Wujudkan Solusi Terbaik Untuk Bisnis Anda"},
			{Name: "subtitle", Label: "Subjudul", Type: FTextarea, Required: true, Full: true, Default: "Temukan solusi profesional dan terintegrasi berkualitas tinggi untuk mendukung pertumbuhan bisnis Anda."},
			{Name: "primary_btn_text", Label: "Teks Tombol Utama", Type: FText, Default: "Mulai Proyek"},
			{Name: "primary_btn_url", Label: "Tujuan Tombol Utama", Type: FSelect, Options: btnURLOptions, Default: "whatsapp"},
			{Name: "secondary_btn_text", Label: "Teks Tombol Sekunder", Type: FText, Default: "Layanan Kami"},
			{Name: "secondary_btn_url", Label: "Tujuan Tombol Sekunder", Type: FSelect, Options: btnURLOptions, Default: "#services"},
			{Name: "image_path", Label: "Hero Banner Image", Type: FImage, Dir: "hero"},
		},
		Columns: []Column{{Name: "title"}},
	},
	{
		Group: "Hero", Slug: "hero-backgrounds", Label: "Hero Background", Singular: "Hero Background", Icon: "wallpaper",
		Table: "hero_backgrounds", Singleton: true,
		Fields: []Field{
			{Name: "badge", Label: "Lencana / Teks Kecil", Type: FText},
			{Name: "title", Label: "Judul Utama", Type: FText, Required: true, Default: "Wujudkan Solusi Terbaik Untuk Bisnis Anda"},
			{Name: "subtitle", Label: "Subjudul", Type: FTextarea, Required: true, Full: true},
			{Name: "primary_btn_text", Label: "Teks Tombol Utama", Type: FText, Default: "Mulai Proyek"},
			{Name: "primary_btn_url", Label: "Tujuan Tombol Utama", Type: FSelect, Options: btnURLOptions, Default: "whatsapp"},
			{Name: "secondary_btn_text", Label: "Teks Tombol Sekunder", Type: FText, Default: "Layanan Kami"},
			{Name: "secondary_btn_url", Label: "Tujuan Tombol Sekunder", Type: FSelect, Options: btnURLOptions, Default: "#services"},
			{Name: "image_path", Label: "Background Image", Type: FImage, Dir: "hero_background"},
		},
		Columns: []Column{{Name: "title"}},
	},
	{
		Group: "Hero", Slug: "hero-carousels", Label: "Hero Carousel", Singular: "Hero Carousel", Icon: "images",
		Table: "hero_carousels", Singleton: true,
		Fields: []Field{
			{Name: "badge", Label: "Lencana / Teks Kecil", Type: FText},
			{Name: "title", Label: "Judul Utama", Type: FText, Required: true, Default: "Wujudkan Solusi Terbaik Untuk Bisnis Anda"},
			{Name: "subtitle", Label: "Subjudul", Type: FTextarea, Required: true, Full: true},
			{Name: "primary_btn_text", Label: "Teks Tombol Utama", Type: FText, Default: "Mulai Proyek"},
			{Name: "primary_btn_url", Label: "Tujuan Tombol Utama", Type: FSelect, Options: btnURLOptions, Default: "whatsapp"},
			{Name: "secondary_btn_text", Label: "Teks Tombol Sekunder", Type: FText, Default: "Layanan Kami"},
			{Name: "secondary_btn_url", Label: "Tujuan Tombol Sekunder", Type: FSelect, Options: btnURLOptions, Default: "#services"},
		},
		Columns: []Column{{Name: "title"}},
	},

	// ----------------- PROFIL PERUSAHAAN -----------------
	{
		Group: "Profil Perusahaan", Slug: "about-items", Label: "About Items", Singular: "About Item", Icon: "info",
		Table: "about_items", DefaultSort: "sort_order",
		Fields: []Field{
			{Name: "title", Label: "Judul / Nilai Unggulan", Type: FText, Required: true},
			{Name: "description", Label: "Deskripsi", Type: FTextarea, Full: true},
			{Name: "image_path", Label: "Icon / Ilustrasi", Type: FImage, Dir: "about_items"},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "title", Searchable: true},
			{Name: "sort_order", Sortable: true},
		},
	},
	{
		Group: "Profil Perusahaan", Slug: "work-steps", Label: "Cara Kerja (Work Steps)", Singular: "Work Step", Icon: "footprints",
		Table: "work_steps", DefaultSort: "sort_order",
		Fields: []Field{
			{Name: "title", Label: "Judul Langkah", Type: FText, Required: true},
			{Name: "description", Label: "Deskripsi Langkah", Type: FTextarea, Full: true},
			{Name: "image_path", Label: "Ilustrasi", Type: FImage, Dir: "work_steps"},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "title", Searchable: true},
			{Name: "sort_order", Sortable: true},
		},
	},
	{
		Group: "Profil Perusahaan", Slug: "why-choose-us", Label: "Keunggulan (Why Us)", Singular: "Why Choose Us", Icon: "sparkles",
		Table: "why_choose_us", Singleton: true,
		Fields: []Field{
			{Name: "title", Label: "Judul Utama", Type: FText, Required: true, Default: "Mengapa Memilih Kami?"},
			{Name: "subtitle", Label: "Subjudul", Type: FTextarea, Required: true, Full: true},
			{Name: "features", Label: "Checklist Benefits (Tekan Enter)", Type: FTags, Required: true, Full: true, Placeholder: "Tambah keunggulan...",
				DefaultTags: []string{"Solusi terintegrasi dan berkualitas tinggi", "Tim profesional dan berpengalaman", "Pengerjaan tepat waktu", "Dukungan teknis responsif"}},
			{Name: "image_path", Label: "Side Mockup Image", Type: FImage, Dir: "why_choose_us"},
		},
		Columns: []Column{{Name: "title"}},
	},
	{
		Group: "Profil Perusahaan", Slug: "branches", Label: "Cabang & Peta", Singular: "Cabang", Icon: "map-pin",
		Table: "branches", DefaultSort: "sort_order",
		Fields: []Field{
			{Name: "name", Label: "Nama Cabang", Type: FText, Required: true},
			{Name: "phone", Label: "Nomor Telepon/WA", Type: FText},
			{Name: "address", Label: "Alamat Lengkap", Type: FTextarea, Required: true, Full: true},
			{Name: "latitude", Label: "Latitude", Type: FText, Required: true, Placeholder: "-6.208763"},
			{Name: "longitude", Label: "Longitude", Type: FText, Required: true, Placeholder: "106.845599"},
			{Name: "is_active", Label: "Aktif", Type: FToggle, Default: "true"},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "name", Searchable: true},
			{Name: "phone", Searchable: true},
			{Name: "is_active", Kind: CBool, Sortable: true},
			{Name: "sort_order", Sortable: true},
		},
	},

	// ----------------- LAYANAN & PORTOFOLIO -----------------
	{
		Group: "Layanan & Portofolio", Slug: "services", Label: "Layanan Utama", Singular: "Layanan", Icon: "layers",
		Table: "services", DefaultSort: "sort_order",
		Fields: []Field{
			{Name: "title", Label: "Judul Layanan", Type: FText, Required: true},
			{Name: "description", Label: "Deskripsi", Type: FTextarea, Required: true, Full: true},
			{Name: "icon", Label: "Icon Lucide", Type: FText, Required: true, Placeholder: "globe, layout, cpu, printer, package"},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "title", Searchable: true},
			{Name: "icon", Searchable: true},
			{Name: "sort_order", Sortable: true},
		},
	},
	{
		Group: "Layanan & Portofolio", Slug: "service-premiums", Label: "Layanan Premium", Singular: "Layanan Premium", Icon: "gem",
		Table: "service_premiums", DefaultSort: "sort_order",
		Fields: []Field{
			{Name: "title", Label: "Nama Layanan", Type: FText, Required: true},
			{Name: "description", Label: "Deskripsi", Type: FTextarea, Full: true},
			{Name: "image_path", Label: "Gambar / Line Art", Type: FImage, Dir: "service_premiums"},
			{Name: "button_text", Label: "Teks Tombol", Type: FText, Default: "Start Project", Required: true},
			{Name: "button_url", Label: "URL Tombol", Type: FText, Default: "whatsapp", Required: true},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "title", Searchable: true},
			{Name: "button_text"},
			{Name: "sort_order", Sortable: true},
		},
	},
	{
		Group: "Layanan & Portofolio", Slug: "portfolios", Label: "Portofolio", Singular: "Portofolio", Icon: "image",
		Table: "portfolios", DefaultSort: "sort_order",
		Fields: []Field{
			{Name: "title", Label: "Judul Proyek", Type: FText, Required: true},
			{Name: "category", Label: "Kategori", Type: FText, Required: true, Placeholder: "Website, Mobile App, Branding"},
			{Name: "image_path", Label: "Gambar Portofolio", Type: FImage, Required: true, Dir: "portfolios"},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "image_path", Label: "Image", Kind: CImage},
			{Name: "title", Searchable: true},
			{Name: "category", Searchable: true, Sortable: true},
			{Name: "sort_order", Sortable: true},
		},
	},
	{
		Group: "Layanan & Portofolio", Slug: "order-steps", Label: "Langkah Order", Singular: "Langkah Order", Icon: "list-ordered",
		Table: "order_steps", DefaultSort: "step_number",
		Fields: []Field{
			{Name: "step_number", Label: "Nomor Langkah", Type: FNumber, Required: true},
			{Name: "title", Label: "Judul", Type: FText, Required: true},
			{Name: "description", Label: "Deskripsi", Type: FTextarea, Required: true, Full: true},
			{Name: "icon", Label: "Icon Lucide", Type: FText, Placeholder: "message-square, edit, briefcase, check"},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "step_number", Sortable: true},
			{Name: "title", Searchable: true},
			{Name: "sort_order", Sortable: true},
		},
	},

	// ----------------- TESTIMONI & KLIEN -----------------
	{
		Group: "Testimoni & Klien", Slug: "testimonials", Label: "Testimoni", Singular: "Testimoni", Icon: "message-square-quote",
		Table: "testimonials", DefaultSort: "sort_order",
		Fields: []Field{
			{Name: "client_name", Label: "Nama Klien", Type: FText, Required: true},
			{Name: "client_role", Label: "Peran / Perusahaan", Type: FText},
			{Name: "stars", Label: "Bintang", Type: FSelect, Required: true, Default: "5", Options: []Option{
				{"5", "★★★★★ (5 Stars)"}, {"4", "★★★★ (4 Stars)"}, {"3", "★★★ (3 Stars)"}, {"2", "★★ (2 Stars)"}, {"1", "★ (1 Star)"},
			}},
			{Name: "content", Label: "Isi Testimoni", Type: FTextarea, Required: true, Full: true},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "client_name", Searchable: true},
			{Name: "client_role", Searchable: true},
			{Name: "stars", Kind: CStars, Sortable: true},
			{Name: "sort_order", Sortable: true},
		},
	},
	{
		Group: "Testimoni & Klien", Slug: "clients", Label: "Klien & Mitra", Singular: "Klien", Icon: "handshake",
		Table: "clients", DefaultSort: "sort_order",
		Fields: []Field{
			{Name: "name", Label: "Nama Klien / Perusahaan", Type: FText, Required: true},
			{Name: "logo", Label: "Logo Klien", Type: FImage, Required: true, Dir: "clients"},
			{Name: "website_url", Label: "Website URL", Type: FURL},
			{Name: "is_active", Label: "Aktif", Type: FToggle, Default: "true"},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "logo", Label: "Logo", Kind: CImage},
			{Name: "name", Searchable: true},
			{Name: "is_active", Kind: CBool, Sortable: true},
			{Name: "sort_order", Sortable: true},
		},
	},
	{
		Group: "Testimoni & Klien", Slug: "statistics", Label: "Statistik", Singular: "Statistik", Icon: "chart-column",
		Table: "statistics", DefaultSort: "sort_order",
		Fields: []Field{
			{Name: "value", Label: "Nilai", Type: FText, Required: true, Placeholder: "500+, 1,200+, 5 Tahun"},
			{Name: "label", Label: "Keterangan", Type: FText, Required: true, Placeholder: "Klien Puas, Proyek Selesai"},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "value", Searchable: true},
			{Name: "label", Searchable: true},
			{Name: "sort_order", Sortable: true},
		},
	},

	// ----------------- KONTEN & PROMOSI -----------------
	{
		Group: "Konten & Promosi", Slug: "news", Label: "Berita & News", Singular: "Berita", Icon: "newspaper",
		Table: "news", DefaultSort: "created_at",
		Fields: []Field{
			{Name: "title", Label: "Judul Berita", Type: FText, Required: true},
			{Name: "slug", Label: "Slug (URL)", Type: FText, Placeholder: "dibuat otomatis jika kosong"},
			{Name: "thumbnail", Label: "Thumbnail", Type: FImage, Dir: "news-thumbnails"},
			{Name: "description", Label: "Isi Berita", Type: FTextarea, Full: true},
		},
		Columns: []Column{
			{Name: "thumbnail", Label: "Thumbnail", Kind: CImage},
			{Name: "title", Searchable: true},
			{Name: "slug", Searchable: true},
		},
	},
	{
		Group: "Konten & Promosi", Slug: "activities", Label: "Kegiatan & Acara", Singular: "Kegiatan", Icon: "calendar",
		Table: "activities", DefaultSort: "created_at",
		Fields: []Field{
			{Name: "title", Label: "Judul Kegiatan", Type: FText, Required: true},
			{Name: "slug", Label: "Slug (URL)", Type: FText, Placeholder: "dibuat otomatis jika kosong"},
			{Name: "thumbnail", Label: "Thumbnail", Type: FImage, Dir: "activities-thumbnails"},
			{Name: "description", Label: "Deskripsi", Type: FTextarea, Full: true},
		},
		Columns: []Column{
			{Name: "thumbnail", Label: "Thumbnail", Kind: CImage},
			{Name: "title", Searchable: true},
			{Name: "slug", Searchable: true},
		},
	},
	{
		Group: "Konten & Promosi", Slug: "cta-sections", Label: "CTA Banner", Singular: "CTA Banner", Icon: "megaphone",
		Table: "cta_sections", Singleton: true,
		Fields: []Field{
			{Name: "title", Label: "Judul Banner", Type: FText, Required: true, Default: "Siap Mewujudkan Proyek Anda?"},
			{Name: "subtitle", Label: "Subjudul", Type: FTextarea, Required: true, Full: true},
			{Name: "btn_text", Label: "Teks Tombol", Type: FText, Default: "Mulai Konsultasi"},
			{Name: "btn_url", Label: "Tujuan Tombol", Type: FSelect, Options: btnURLOptions, Default: "whatsapp"},
		},
		Columns: []Column{{Name: "title"}},
	},
	{
		Group: "Konten & Promosi", Slug: "leads", Label: "Pesan Masuk (Leads)", Singular: "Pesan Masuk", Icon: "inbox",
		Table: "leads", DefaultSort: "created_at", ReadOnly: true,
		Fields: []Field{
			{Name: "name", Label: "Nama Pengirim", Type: FReadOnly},
			{Name: "phone", Label: "Nomor WhatsApp / HP", Type: FReadOnly},
			{Name: "message", Label: "Isi Pesan", Type: FReadOnly, Full: true},
		},
		Columns: []Column{
			{Name: "name", Searchable: true},
			{Name: "phone", Searchable: true},
			{Name: "created_at", Label: "Tanggal", Sortable: true},
		},
	},

	// ----------------- PENGATURAN -----------------
	{
		Group: "Pengaturan", Slug: "site-configs", Label: "Pengaturan Umum", Singular: "Pengaturan Umum", Icon: "settings",
		Table: "site_configs", Singleton: true,
		Fields: []Field{
			{Name: "site_name", Label: "Nama Situs / Bisnis", Type: FText, Required: true, Default: "FourPlusOne"},
			{Name: "logo", Label: "Logo Bisnis", Type: FImage, Dir: "logos"},
			{Name: "whatsapp_number", Label: "Nomor WhatsApp (Format: 628...)", Type: FText, Required: true, Placeholder: "628123456789"},
			{Name: "email", Label: "Email", Type: FEmail},
			{Name: "address", Label: "Alamat Kantor", Type: FTextarea, Full: true},
			{Name: "facebook_url", Label: "Facebook URL", Type: FURL},
			{Name: "instagram_url", Label: "Instagram URL", Type: FURL},
			{Name: "twitter_url", Label: "Twitter URL", Type: FURL},
			{Name: "primary_color", Label: "Warna Utama", Type: FText, Default: "#059669"},
			{Name: "secondary_color", Label: "Warna Sekunder", Type: FText, Default: "#0F172A"},
			{Name: "accent_color", Label: "Warna Aksen", Type: FText, Default: "#10B981"},
			{Name: "background_color", Label: "Warna Latar", Type: FText, Default: "#F8FAFC"},
			{Name: "contact_title", Label: "Judul Kontak", Type: FText},
			{Name: "contact_subtitle", Label: "Subjudul Kontak", Type: FTextarea, Full: true},
			{Name: "news_title", Label: "Judul Berita", Type: FText},
			{Name: "news_subtitle", Label: "Subjudul Berita", Type: FTextarea, Full: true},
			{Name: "activities_title", Label: "Judul Kegiatan", Type: FText},
			{Name: "activities_subtitle", Label: "Subjudul Kegiatan", Type: FTextarea, Full: true},
			{Name: "services_title", Label: "Judul Layanan", Type: FText},
			{Name: "services_subtitle", Label: "Subjudul Layanan", Type: FTextarea, Full: true},
			{Name: "portfolio_title", Label: "Judul Portofolio", Type: FText},
			{Name: "portfolio_subtitle", Label: "Subjudul Portofolio", Type: FTextarea, Full: true},
			{Name: "order_steps_title", Label: "Judul Cara Pesan", Type: FText},
			{Name: "order_steps_subtitle", Label: "Subjudul Cara Pesan", Type: FTextarea, Full: true},
			{Name: "testimonials_title", Label: "Judul Testimoni", Type: FText},
			{Name: "testimonials_subtitle", Label: "Subjudul Testimoni", Type: FTextarea, Full: true},
			{Name: "footer_description", Label: "Deskripsi Footer", Type: FTextarea, Full: true},
			{Name: "about_title", Label: "Judul About", Type: FText},
			{Name: "about_subtitle", Label: "Subjudul About", Type: FTextarea, Full: true},
			{Name: "service_premium_title", Label: "Judul Layanan Premium", Type: FText},
			{Name: "service_premium_subtitle", Label: "Subjudul Layanan Premium", Type: FTextarea, Full: true},
			{Name: "work_steps_title", Label: "Judul Cara Kerja", Type: FText},
			{Name: "work_steps_subtitle", Label: "Subjudul Cara Kerja", Type: FTextarea, Full: true},
			{Name: "branches_badge", Label: "Badge Cabang", Type: FText},
			{Name: "branches_title", Label: "Judul Cabang", Type: FText},
			{Name: "branches_subtitle", Label: "Subjudul Cabang", Type: FTextarea, Full: true},
			{Name: "clients_title", Label: "Judul Klien", Type: FText},
			{Name: "clients_subtitle", Label: "Subjudul Klien", Type: FTextarea, Full: true},
		},
		Columns: []Column{{Name: "site_name"}, {Name: "whatsapp_number"}},
	},
	{
		Group: "Pengaturan", Slug: "landing-sections", Label: "Visibilitas Bagian (Sections)", Singular: "Bagian", Icon: "toggle-left",
		Table: "landing_sections", DefaultSort: "sort_order",
		Fields: []Field{
			{Name: "section_key", Label: "Kunci Bagian", Type: FText, Required: true},
			{Name: "title", Label: "Nama Bagian", Type: FText, Required: true},
			{Name: "is_active", Label: "Tampilkan di Landing Page", Type: FToggle, Default: "true"},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "section_key"},
			{Name: "title", Searchable: true},
			{Name: "is_active", Kind: CBool, Sortable: true},
			{Name: "sort_order", Sortable: true},
		},
	},

	// ----------------- MANAJEMEN PENGGUNA -----------------
	{
		Group: "Manajemen Pengguna", Slug: "users", Label: "Pengguna Admin", Singular: "Pengguna", Icon: "users",
		Table: "users", DefaultSort: "name",
		Fields: []Field{
			{Name: "name", Label: "Nama Lengkap", Type: FText, Required: true},
			{Name: "email", Label: "Alamat Email", Type: FEmail, Required: true},
		},
		Columns: []Column{
			{Name: "name", Searchable: true},
			{Name: "email", Searchable: true},
		},
	},
}

type GroupedResources struct {
	Group string
	Items []*Resource
}

func GroupedNav() []GroupedResources {
	seen := map[string]*GroupedResources{}
	var order []string
	for _, r := range Resources {
		g := r.Group
		if g == "" {
			g = "Lainnya"
		}
		if _, ok := seen[g]; !ok {
			seen[g] = &GroupedResources{Group: g}
			order = append(order, g)
		}
		seen[g].Items = append(seen[g].Items, r)
	}
	var res []GroupedResources
	for _, g := range order {
		res = append(res, *seen[g])
	}
	return res
}

func findResource(slug string) *Resource {
	for _, r := range Resources {
		if r.Slug == slug {
			return r
		}
	}
	return nil
}
