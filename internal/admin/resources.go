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
	FRepeater FieldType = "repeater"
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
	Hint        string
	Sub         []Field // repeater item fields
	AddLabel    string
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
	CToggle ColKind = "toggle"
	CBadge  ColKind = "badge"
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
	NoCreate    bool // fixed records (e.g. landing-sections)
	Reorder     bool // drag & drop ordering by sort_order
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
	return Field{Name: "sort_order", Label: "Urutan Tampilan", Type: FNumber, Required: true, Default: "0"}
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

// Resources is the full panel definition for branch General (based on BE LP Filament).
var Resources = []*Resource{
	// ----------------- UNGROUPED / TOP LEVEL -----------------
	{
		Group: "", Slug: "leads", Label: "Leads / Pesan", Singular: "Lead", Icon: "inbox",
		Table: "leads", DefaultSort: "created_at", ReadOnly: true,
		Fields: []Field{
			{Name: "name", Label: "Nama Pengirim", Type: FReadOnly},
			{Name: "phone", Label: "Nomor Telepon/WA", Type: FReadOnly},
			{Name: "message", Label: "Pesan", Type: FReadOnly, Full: true},
		},
		Columns: []Column{
			{Name: "name", Label: "Nama Pengirim", Searchable: true},
			{Name: "phone", Label: "Nomor WA", Searchable: true},
			{Name: "message", Label: "Pesan", Searchable: true},
			{Name: "created_at", Label: "Tanggal Masuk", Sortable: true},
		},
	},

	// ----------------- HERO -----------------
	{
		Group: "Hero", Slug: "hero-sections", Label: "Basic Hero", Singular: "Basic Hero", Icon: "house",
		Table: "hero_sections", Singleton: true,
		Fields: []Field{
			{Name: "badge", Label: "Lencana / Teks Kecil", Type: FText, Placeholder: "contoh: Bisnis Kuliner Terbaik"},
			{Name: "title", Label: "Judul Utama", Type: FText, Required: true, Hint: "Gunakan kurung siku [...] untuk mewarnai kata tertentu (contoh: Wujudkan Ide Anda [Dalam Cetakan]).", Default: "Wujudkan Ide Anda Dalam Cetakan"},
			{Name: "subtitle", Label: "Subjudul", Type: FTextarea, Required: true, Full: true, Hint: "Gunakan kurung siku [...] untuk mewarnai kata tertentu.", Default: "Temukan solusi percetakan digital berkualitas terbaik untuk spanduk, brosur, stiker, dan kemasan Anda."},
			{Name: "primary_btn_text", Label: "Teks Tombol Utama", Type: FText, Default: "Pesan Sekarang"},
			{Name: "primary_btn_url", Label: "Tujuan Tombol Utama", Type: FSelect, Required: true, Options: btnURLOptions, Default: "whatsapp"},
			{Name: "secondary_btn_text", Label: "Teks Tombol Sekunder", Type: FText, Default: "Layanan Kami"},
			{Name: "secondary_btn_url", Label: "Tujuan Tombol Sekunder", Type: FSelect, Required: true, Options: btnURLOptions, Default: "#services"},
			{Name: "image_path", Label: "Gambar Banner Hero", Type: FImage, Dir: "hero"},
		},
		Columns: []Column{{Name: "title", Label: "Judul"}},
	},
	{
		Group: "Hero", Slug: "hero-backgrounds", Label: "Hero with Background", Singular: "Hero with Background", Icon: "wallpaper",
		Table: "hero_backgrounds", Singleton: true,
		Fields: []Field{
			{Name: "badge", Label: "Lencana / Teks Kecil", Type: FText, Placeholder: "contoh: Bisnis Kuliner Terbaik"},
			{Name: "title", Label: "Judul Utama", Type: FText, Required: true, Hint: "Gunakan kurung siku [...] untuk mewarnai kata tertentu (contoh: Wujudkan Ide Anda [Dalam Cetakan]).", Default: "Wujudkan Ide Anda Dalam Cetakan"},
			{Name: "subtitle", Label: "Subjudul", Type: FTextarea, Required: true, Full: true, Hint: "Gunakan kurung siku [...] untuk mewarnai kata tertentu.", Default: "Temukan solusi percetakan digital berkualitas terbaik untuk spanduk, brosur, stiker, dan kemasan Anda."},
			{Name: "primary_btn_text", Label: "Teks Tombol Utama", Type: FText, Default: "Pesan Sekarang"},
			{Name: "primary_btn_url", Label: "Tujuan Tombol Utama", Type: FSelect, Required: true, Options: btnURLOptions, Default: "whatsapp"},
			{Name: "secondary_btn_text", Label: "Teks Tombol Sekunder", Type: FText, Default: "Layanan Kami"},
			{Name: "secondary_btn_url", Label: "Tujuan Tombol Sekunder", Type: FSelect, Required: true, Options: btnURLOptions, Default: "#services"},
			{Name: "image_path", Label: "Gambar Latar Belakang (Background Image)", Type: FImage, Dir: "hero_background"},
		},
		Columns: []Column{{Name: "title", Label: "Judul"}},
	},
	{
		Group: "Hero", Slug: "hero-carousels", Label: "Hero with Carousel", Singular: "Hero with Carousel", Icon: "images",
		Table: "hero_carousels", Singleton: true,
		Fields: []Field{
			{Name: "badge", Label: "Lencana / Teks Kecil", Type: FText, Placeholder: "contoh: Bisnis Kuliner Terbaik"},
			{Name: "title", Label: "Judul Utama", Type: FText, Required: true, Hint: "Gunakan kurung siku [...] untuk mewarnai kata tertentu (contoh: Wujudkan Ide Anda [Dalam Cetakan]).", Default: "Wujudkan Ide Anda Dalam Cetakan"},
			{Name: "subtitle", Label: "Subjudul", Type: FTextarea, Required: true, Full: true, Hint: "Gunakan kurung siku [...] untuk mewarnai kata tertentu.", Default: "Temukan solusi percetakan digital berkualitas terbaik untuk spanduk, brosur, stiker, dan kemasan Anda."},
			{Name: "primary_btn_text", Label: "Teks Tombol Utama", Type: FText, Default: "Pesan Sekarang"},
			{Name: "primary_btn_url", Label: "Tujuan Tombol Utama", Type: FSelect, Required: true, Options: btnURLOptions, Default: "whatsapp"},
			{Name: "secondary_btn_text", Label: "Teks Tombol Sekunder", Type: FText, Default: "Layanan Kami"},
			{Name: "secondary_btn_url", Label: "Tujuan Tombol Sekunder", Type: FSelect, Required: true, Options: btnURLOptions, Default: "#services"},
			{Name: "carousel_images", Label: "Item Carousel / Produk Unggulan", Type: FRepeater, Full: true, AddLabel: "Add to item Carousel / Produk Unggulan", Sub: []Field{
				{Name: "image", Label: "Gambar Produk", Type: FImage, Required: true, Dir: "hero_carousel"},
				{Name: "title", Label: "Nama Produk / Label Utama", Type: FText, Required: true, Placeholder: "contoh: ENGLISH BREAKFAST"},
				{Name: "subtitle", Label: "Keterangan / Sub-label", Type: FText, Placeholder: "contoh: Strong & Malty"},
				{Name: "footer", Label: "Keterangan Kaki (Harga/Bintang)", Type: FText, Placeholder: "contoh: dari Rp 50.000 | ★★★★★"},
			}},
		},
		Columns: []Column{{Name: "title", Label: "Judul"}},
	},

	// ----------------- PROFIL PERUSAHAAN -----------------
	{
		Group: "Profil Perusahaan", Slug: "about-items", Label: "Tentang Kami", Singular: "Tentang Kami", Icon: "info",
		Table: "about_items", DefaultSort: "sort_order", Reorder: true,
		Fields: []Field{
			{Name: "title", Label: "Judul", Type: FText, Required: true},
			{Name: "description", Label: "Deskripsi", Type: FTextarea, Full: true},
			{Name: "image_path", Label: "Gambar", Type: FImage, Dir: "about_items"},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "image_path", Label: "Gambar", Kind: CImage},
			{Name: "title", Label: "Judul", Searchable: true},
			{Name: "description", Label: "Deskripsi"},
			{Name: "sort_order", Label: "Urutan", Sortable: true},
		},
	},
	{
		Group: "Profil Perusahaan", Slug: "why-choose-us", Label: "Keunggulan Kami", Singular: "Keunggulan", Icon: "sparkles",
		Table: "why_choose_us", Singleton: true,
		Fields: []Field{
			{Name: "title", Label: "Judul", Type: FText, Required: true, Default: "Mengapa Memilih Kami?"},
			{Name: "subtitle", Label: "Subjudul", Type: FTextarea, Required: true, Full: true},
			{Name: "features", Label: "Checklist Benefits (Tekan Enter)", Type: FTags, Required: true, Full: true, Placeholder: "Tambah keunggulan...",
				DefaultTags: []string{"Solusi terintegrasi dan berkualitas tinggi", "Tim profesional dan berpengalaman", "Pengerjaan tepat waktu", "Dukungan teknis responsif"}},
			{Name: "image_path", Label: "Gambar Mockup Samping", Type: FImage, Dir: "why_choose_us"},
		},
		Columns: []Column{{Name: "title", Label: "Judul"}},
	},
	{
		Group: "Profil Perusahaan", Slug: "statistics", Label: "Statistik", Singular: "Statistik", Icon: "chart-column",
		Table: "statistics", DefaultSort: "sort_order", Reorder: true,
		Fields: []Field{
			{Name: "value", Label: "Nilai / Angka", Type: FText, Required: true, Placeholder: "contoh: 500+, 1,200+, 5 Tahun"},
			{Name: "label", Label: "Label / Nama Statistik", Type: FText, Required: true, Placeholder: "contoh: Klien Puas, Proyek Selesai"},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "value", Label: "Nilai", Searchable: true},
			{Name: "label", Label: "Nama Statistik", Searchable: true},
			{Name: "sort_order", Label: "Urutan", Sortable: true},
		},
	},

	// ----------------- LAYANAN & PORTOFOLIO -----------------
	{
		Group: "Layanan & Portofolio", Slug: "services", Label: "Layanan", Singular: "Layanan", Icon: "layers",
		Table: "services", DefaultSort: "sort_order", Reorder: true,
		Fields: []Field{
			{Name: "title", Label: "Nama Layanan", Type: FText, Required: true},
			{Name: "description", Label: "Deskripsi Layanan", Type: FTextarea, Required: true, Full: true},
			{Name: "icon", Label: "Ikon (Lucide)", Type: FText, Required: true, Placeholder: "globe, layout, cpu, printer, package"},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "title", Label: "Nama Layanan", Searchable: true},
			{Name: "icon", Label: "Ikon", Searchable: true},
			{Name: "sort_order", Label: "Urutan", Sortable: true},
		},
	},
	{
		Group: "Layanan & Portofolio", Slug: "service-premiums", Label: "Layanan Premium", Singular: "Layanan Premium", Icon: "gem",
		Table: "service_premiums", DefaultSort: "sort_order", Reorder: true,
		Fields: []Field{
			{Name: "title", Label: "Nama Layanan", Type: FText, Required: true},
			{Name: "description", Label: "Deskripsi Layanan", Type: FTextarea, Full: true},
			{Name: "image_path", Label: "Gambar / Ilustrasi Line Art", Type: FImage, Dir: "service_premiums"},
			{Name: "button_text", Label: "Teks Tombol", Type: FText, Default: "Start Project", Required: true},
			{Name: "button_url", Label: "URL Tombol", Type: FText, Default: "whatsapp", Required: true},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "image_path", Label: "Gambar/Ilustrasi", Kind: CImage},
			{Name: "title", Label: "Nama Layanan", Searchable: true},
			{Name: "description", Label: "Deskripsi"},
			{Name: "button_text", Label: "Teks Tombol"},
			{Name: "sort_order", Label: "Urutan", Sortable: true},
		},
	},
	{
		Group: "Layanan & Portofolio", Slug: "portfolios", Label: "Portofolio", Singular: "Portofolio", Icon: "image",
		Table: "portfolios", DefaultSort: "sort_order", Reorder: true,
		Fields: []Field{
			{Name: "title", Label: "Nama Portofolio", Type: FText, Required: true},
			{Name: "category", Label: "Kategori", Type: FText, Required: true, Placeholder: "Website, Mobile App, Branding"},
			{Name: "image_path", Label: "Gambar Portofolio", Type: FImage, Required: true, Dir: "portfolios"},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "image_path", Label: "Gambar", Kind: CImage},
			{Name: "title", Label: "Nama Portofolio", Searchable: true},
			{Name: "category", Label: "Kategori", Searchable: true, Sortable: true},
			{Name: "sort_order", Label: "Urutan", Sortable: true},
		},
	},
	{
		Group: "Layanan & Portofolio", Slug: "order-steps", Label: "Langkah Pemesanan", Singular: "Langkah Pemesanan", Icon: "list-ordered",
		Table: "order_steps", DefaultSort: "step_number", Reorder: true,
		Fields: []Field{
			{Name: "step_number", Label: "Nomor Langkah", Type: FNumber, Required: true},
			{Name: "title", Label: "Judul Langkah", Type: FText, Required: true},
			{Name: "description", Label: "Deskripsi Langkah", Type: FTextarea, Required: true, Full: true},
			{Name: "icon", Label: "Ikon (Lucide)", Type: FText, Placeholder: "message-square, edit, briefcase, check"},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "step_number", Label: "Nomor Langkah", Sortable: true},
			{Name: "title", Label: "Judul Langkah", Searchable: true},
			{Name: "icon", Label: "Ikon"},
			{Name: "sort_order", Label: "Urutan", Sortable: true},
		},
	},
	{
		Group: "Layanan & Portofolio", Slug: "work-steps", Label: "Cara Kerja", Singular: "Langkah Kerja", Icon: "footprints",
		Table: "work_steps", DefaultSort: "sort_order", Reorder: true,
		Fields: []Field{
			{Name: "title", Label: "Judul Langkah", Type: FText, Required: true},
			{Name: "description", Label: "Deskripsi Langkah", Type: FTextarea, Full: true},
			{Name: "image_path", Label: "Ilustrasi Lingkaran (Circular Illustration)", Type: FImage, Dir: "work_steps"},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "image_path", Label: "Ilustrasi", Kind: CImage},
			{Name: "title", Label: "Judul Langkah", Searchable: true},
			{Name: "description", Label: "Deskripsi"},
			{Name: "sort_order", Label: "Urutan", Sortable: true},
		},
	},

	// ----------------- TESTIMONI & KLIEN -----------------
	{
		Group: "Testimoni & Klien", Slug: "testimonials", Label: "Testimoni", Singular: "Testimoni", Icon: "message-square-quote",
		Table: "testimonials", DefaultSort: "sort_order", Reorder: true,
		Fields: []Field{
			{Name: "client_name", Label: "Nama Klien", Type: FText, Required: true},
			{Name: "client_role", Label: "Pekerjaan / Jabatan", Type: FText},
			{Name: "stars", Label: "Bintang Penilaian", Type: FSelect, Required: true, Default: "5", Options: []Option{
				{"5", "★★★★★ (5 Stars)"}, {"4", "★★★★ (4 Stars)"}, {"3", "★★★ (3 Stars)"}, {"2", "★★ (2 Stars)"}, {"1", "★ (1 Star)"},
			}},
			{Name: "content", Label: "Isi Testimoni", Type: FTextarea, Required: true, Full: true},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "client_name", Label: "Nama Klien", Searchable: true},
			{Name: "client_role", Label: "Pekerjaan / Jabatan", Searchable: true},
			{Name: "stars", Label: "Bintang", Kind: CStars, Sortable: true},
			{Name: "sort_order", Label: "Urutan", Sortable: true},
		},
	},
	{
		Group: "Testimoni & Klien", Slug: "clients", Label: "Logo Klien / Perusahaan", Singular: "Klien", Icon: "users",
		Table: "clients", DefaultSort: "sort_order", Reorder: true,
		Fields: []Field{
			{Name: "name", Label: "Nama Klien / Perusahaan", Type: FText, Required: true},
			{Name: "logo", Label: "Logo Klien", Type: FImage, Required: true, Dir: "clients"},
			{Name: "website_url", Label: "Tautan Website Klien (Opsional)", Type: FURL},
			{Name: "is_active", Label: "Aktif", Type: FToggle, Default: "true"},
			sortOrderField(),
		},
		Columns: []Column{
			{Name: "logo", Label: "Logo", Kind: CImage},
			{Name: "name", Label: "Nama Klien / Perusahaan", Searchable: true},
			{Name: "website_url", Label: "Website"},
			{Name: "is_active", Label: "Aktif", Kind: CToggle, Sortable: true},
			{Name: "sort_order", Label: "Urutan", Sortable: true},
		},
	},

	// ----------------- KONTEN & PROMOSI -----------------
	{
		Group: "Konten & Promosi", Slug: "news", Label: "Berita", Singular: "Berita", Icon: "newspaper",
		Table: "news", DefaultSort: "created_at",
		Fields: []Field{
			{Name: "title", Label: "Judul Berita", Type: FText, Required: true},
			{Name: "slug", Label: "Slug (URL)", Type: FText, Placeholder: "dibuat otomatis jika kosong"},
			{Name: "thumbnail", Label: "Gambar Thumbnail", Type: FImage, Dir: "news-thumbnails"},
			{Name: "description", Label: "Isi Berita", Type: FTextarea, Full: true},
		},
		Columns: []Column{
			{Name: "thumbnail", Label: "Thumbnail", Kind: CImage},
			{Name: "title", Label: "Judul Berita", Searchable: true, Sortable: true},
			{Name: "created_at", Label: "Tanggal Dibuat", Sortable: true},
		},
	},
	{
		Group: "Konten & Promosi", Slug: "activities", Label: "Aktifitas", Singular: "Aktifitas", Icon: "sparkles",
		Table: "activities", DefaultSort: "created_at",
		Fields: []Field{
			{Name: "title", Label: "Judul Aktifitas", Type: FText, Required: true},
			{Name: "slug", Label: "Slug (URL)", Type: FText, Placeholder: "dibuat otomatis jika kosong"},
			{Name: "thumbnail", Label: "Gambar Thumbnail", Type: FImage, Dir: "activities-thumbnails"},
			{Name: "description", Label: "Deskripsi", Type: FTextarea, Full: true},
		},
		Columns: []Column{
			{Name: "thumbnail", Label: "Thumbnail", Kind: CImage},
			{Name: "title", Label: "Judul Aktifitas", Searchable: true, Sortable: true},
			{Name: "created_at", Label: "Tanggal Dibuat", Sortable: true},
		},
	},
	{
		Group: "Konten & Promosi", Slug: "cta-sections", Label: "Banner CTA", Singular: "Banner CTA", Icon: "megaphone",
		Table: "cta_sections", Singleton: true,
		Fields: []Field{
			{Name: "title", Label: "Judul Banner", Type: FText, Required: true, Default: "Siap Mewujudkan Proyek Anda?"},
			{Name: "subtitle", Label: "Subjudul / Deskripsi Banner", Type: FTextarea, Required: true, Full: true},
			{Name: "btn_text", Label: "Teks Tombol", Type: FText, Default: "Mulai Konsultasi"},
			{Name: "btn_url", Label: "Tujuan Tombol", Type: FSelect, Required: true, Options: btnURLOptions, Default: "whatsapp"},
		},
		Columns: []Column{{Name: "title", Label: "Judul Banner"}},
	},

	// ----------------- PENGATURAN -----------------
	{
		Group: "Pengaturan", Slug: "site-configs", Label: "Pengaturan Umum", Singular: "Pengaturan", Icon: "settings",
		Table: "site_configs", Singleton: true,
		Fields: []Field{
			{Name: "site_name", Label: "Nama Situs / Bisnis", Type: FText, Required: true, Default: "FourPlusOne"},
			{Name: "logo", Label: "Logo Bisnis", Type: FImage, Dir: "logos"},
			{Name: "whatsapp_number", Label: "Nomor WhatsApp (Format: 628...)", Type: FText, Required: true, Placeholder: "628123456789"},
			{Name: "email", Label: "Email", Type: FEmail},
			{Name: "address", Label: "Alamat Kantor / Toko", Type: FTextarea, Full: true},
			{Name: "facebook_url", Label: "URL Facebook", Type: FURL},
			{Name: "instagram_url", Label: "URL Instagram", Type: FURL},
			{Name: "twitter_url", Label: "URL Twitter / X", Type: FURL},
			{Name: "primary_color", Label: "Warna Utama", Type: FText, Default: "#059669"},
			{Name: "secondary_color", Label: "Warna Sekunder", Type: FText, Default: "#0F172A"},
			{Name: "accent_color", Label: "Warna Aksen", Type: FText, Default: "#10B981"},
			{Name: "background_color", Label: "Warna Latar", Type: FText, Default: "#F8FAFC"},
			{Name: "contact_title", Label: "Judul Bagian Kontak", Type: FText},
			{Name: "contact_subtitle", Label: "Deskripsi / Subjudul Bagian Kontak", Type: FTextarea, Full: true},
			{Name: "news_title", Label: "Judul Bagian Berita", Type: FText},
			{Name: "news_subtitle", Label: "Deskripsi / Subjudul Bagian Berita", Type: FTextarea, Full: true},
			{Name: "activities_title", Label: "Judul Bagian Kegiatan", Type: FText},
			{Name: "activities_subtitle", Label: "Deskripsi / Subjudul Bagian Kegiatan", Type: FTextarea, Full: true},
			{Name: "services_title", Label: "Judul Bagian Layanan", Type: FText},
			{Name: "services_subtitle", Label: "Deskripsi / Subjudul Bagian Layanan", Type: FTextarea, Full: true},
			{Name: "portfolio_title", Label: "Judul Bagian Portofolio", Type: FText},
			{Name: "portfolio_subtitle", Label: "Deskripsi / Subjudul Bagian Portofolio", Type: FTextarea, Full: true},
			{Name: "order_steps_title", Label: "Judul Bagian Cara Pesan", Type: FText},
			{Name: "order_steps_subtitle", Label: "Deskripsi / Subjudul Bagian Cara Pesan", Type: FTextarea, Full: true},
			{Name: "testimonials_title", Label: "Judul Bagian Testimoni", Type: FText},
			{Name: "testimonials_subtitle", Label: "Deskripsi / Subjudul Bagian Testimoni", Type: FTextarea, Full: true},
			{Name: "footer_description", Label: "Deskripsi Footer", Type: FTextarea, Full: true},
			{Name: "about_title", Label: "Judul Bagian Tentang Kami", Type: FText},
			{Name: "about_subtitle", Label: "Deskripsi / Subjudul Bagian Tentang Kami", Type: FTextarea, Full: true},
			{Name: "service_premium_title", Label: "Judul Bagian Layanan Premium", Type: FText},
			{Name: "service_premium_subtitle", Label: "Deskripsi / Subjudul Bagian Layanan Premium", Type: FTextarea, Full: true},
			{Name: "work_steps_title", Label: "Judul Bagian Cara Kerja", Type: FText},
			{Name: "work_steps_subtitle", Label: "Deskripsi / Subjudul Bagian Cara Kerja", Type: FTextarea, Full: true},
			{Name: "branches_badge", Label: "Lencana Bagian Cabang (Badge)", Type: FText},
			{Name: "branches_title", Label: "Judul Bagian Cabang", Type: FText},
			{Name: "branches_subtitle", Label: "Deskripsi / Subjudul Bagian Cabang", Type: FTextarea, Full: true},
			{Name: "clients_title", Label: "Judul Bagian Klien", Type: FText},
			{Name: "clients_subtitle", Label: "Deskripsi / Subjudul Bagian Klien", Type: FTextarea, Full: true},
		},
		Columns: []Column{{Name: "site_name", Label: "Nama Situs / Bisnis"}, {Name: "whatsapp_number", Label: "Nomor WhatsApp"}},
	},
	{
		Group: "Pengaturan", Slug: "branches", Label: "Cabang / Lokasi Toko", Singular: "Cabang", Icon: "map-pin",
		Table: "branches", DefaultSort: "sort_order", Reorder: true,
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
			{Name: "name", Label: "Nama Cabang", Searchable: true, Sortable: true},
			{Name: "phone", Label: "Telepon/WA", Searchable: true},
			{Name: "latitude", Label: "Latitude"},
			{Name: "longitude", Label: "Longitude"},
			{Name: "is_active", Label: "Aktif", Kind: CToggle, Sortable: true},
			{Name: "sort_order", Label: "Urutan", Sortable: true},
		},
	},
	{
		Group: "Pengaturan", Slug: "landing-sections", Label: "Urutan & Aktifasi Menu", Singular: "Urutan Menu", Icon: "list-ordered",
		Table: "landing_sections", DefaultSort: "sort_order", NoCreate: true, Reorder: true,
		Fields: []Field{
			{Name: "title", Label: "Nama Menu / Section", Type: FText, Required: true},
			{Name: "is_active", Label: "Aktif", Type: FToggle, Default: "true"},
		},
		Columns: []Column{
			{Name: "title", Label: "Nama Menu / Section", Searchable: true},
			{Name: "section_key", Label: "ID Section", Kind: CBadge, Searchable: true},
			{Name: "is_active", Label: "Status Aktif", Kind: CToggle},
		},
	},

	// ----------------- MANAJEMEN PENGGUNA -----------------
	{
		Group: "Manajemen Pengguna", Slug: "users", Label: "Pengguna", Singular: "Pengguna", Icon: "users",
		Table: "users", DefaultSort: "name",
		Fields: []Field{
			{Name: "name", Label: "Nama", Type: FText, Required: true},
			{Name: "email", Label: "Email", Type: FEmail, Required: true},
		},
		Columns: []Column{
			{Name: "name", Label: "Nama", Searchable: true, Sortable: true},
			{Name: "email", Label: "Email", Searchable: true, Sortable: true},
		},
	},
}

type GroupedResources struct {
	Group string
	Items []*Resource
}

func GroupedNav() []GroupedResources {
	groupOrder := []string{
		"Hero",
		"Profil Perusahaan",
		"Layanan & Portofolio",
		"Testimoni & Klien",
		"Konten & Promosi",
		"Pengaturan",
		"Manajemen Pengguna",
	}
	grouped := map[string]*GroupedResources{}
	for _, g := range groupOrder {
		grouped[g] = &GroupedResources{Group: g}
	}
	for _, r := range Resources {
		if r.Group == "" {
			continue // Ungrouped items like Leads
		}
		if grp, ok := grouped[r.Group]; ok {
			grp.Items = append(grp.Items, r)
		} else {
			grouped[r.Group] = &GroupedResources{Group: r.Group, Items: []*Resource{r}}
			groupOrder = append(groupOrder, r.Group)
		}
	}
	var res []GroupedResources
	for _, g := range groupOrder {
		if grp, ok := grouped[g]; ok && len(grp.Items) > 0 {
			res = append(res, *grp)
		}
	}
	return res
}

func UngroupedNav() []*Resource {
	var out []*Resource
	for _, r := range Resources {
		if r.Group == "" {
			out = append(out, r)
		}
	}
	return out
}

func findResource(slug string) *Resource {
	for _, r := range Resources {
		if r.Slug == slug {
			return r
		}
	}
	return nil
}
