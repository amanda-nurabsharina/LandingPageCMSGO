package seed

import (
	"time"

	"be-lp-go/internal/database"
	"be-lp-go/internal/models"

	"gorm.io/gorm"
)

func now() *time.Time {
	t := time.Now()
	return &t
}

// Singletons returns the single-record configuration rows (id = 1), creating
// them with default values when they do not exist yet.
func Singletons(db *gorm.DB) (*models.SiteConfig, *models.HeroSection, *models.HeroBackground, *models.HeroCarousel, *models.WhyChooseUs, *models.CtaSection, error) {
	site := models.SiteConfig{}
	res := db.Where(models.SiteConfig{ID: 1}).Attrs(models.SiteConfig{
		SiteName:               "FourPlusOne",
		WhatsappNumber:         "628123456789",
		Email:                  models.Str("info@fourplusone.com"),
		Address:                models.Str("Jakarta, Indonesia"),
		FacebookURL:            models.Str("https://facebook.com"),
		InstagramURL:           models.Str("https://instagram.com"),
		TwitterURL:             models.Str("https://twitter.com"),
		PrimaryColor:           models.Str("#059669"),
		SecondaryColor:         models.Str("#0F172A"),
		AccentColor:            models.Str("#10B981"),
		BackgroundColor:        models.Str("#F8FAFC"),
		ContactTitle:           models.Str("Kirimkan Pesan atau Konsultasi Gratis"),
		ContactSubtitle:        models.Str("Punya pertanyaan seputar layanan kami, atau ingin mendiskusikan kebutuhan proyek Anda? Isi formulir, tim kami akan segera menghubungi Anda."),
		NewsTitle:              models.Str("Berita & Informasi Terkini"),
		NewsSubtitle:           models.Str("Ikuti perkembangan terbaru mengenai layanan, inovasi, dan tips seputar industri kami."),
		ActivitiesTitle:        models.Str("Aktifitas & Dokumentasi"),
		ActivitiesSubtitle:     models.Str("Dokumentasi portofolio kerja, aktivitas tim, serta event penting yang kami hadiri."),
		ServicesTitle:          models.Str("Solusi Profesional & Layanan Terintegrasi"),
		ServicesSubtitle:       models.Str("Kami siap menghadirkan berbagai solusi terbaik untuk kebutuhan branding, pengembangan bisnis, dan transformasi digital Anda."),
		PortfolioTitle:         models.Str("Portofolio Proyek Unggulan"),
		PortfolioSubtitle:      models.Str("Berikut adalah beberapa showcase proyek dan solusi terbaik yang telah diselesaikan untuk klien-klien kami yang puas."),
		OrderStepsTitle:        models.Str("Alur Kerja Mudah & Terstruktur"),
		OrderStepsSubtitle:     models.Str("Cukup selesaikan 4 langkah mudah berikut untuk mewujudkan proyek dan kebutuhan bisnis Anda."),
		TestimonialsTitle:      models.Str("Apa Kata Klien & Mitra Kami"),
		TestimonialsSubtitle:   models.Str("Kepercayaan dan kepuasan klien adalah komitmen utama kami. Simak penilaian langsung mereka."),
		FooterDescription:      models.Str("Menyediakan layanan profesional dan solusi digital berkualitas tinggi untuk mendukung pertumbuhan bisnis Anda."),
		AboutTitle:             models.Str("Innovation meets precision."),
		AboutSubtitle:          models.Str("Welcome to Fourplusone. We are a premier IT Software House dedicated to bridging the gap between complex business needs and elegant digital experiences"),
		ServicePremiumTitle:    models.Str("Services We Provide"),
		ServicePremiumSubtitle: models.Str("Tailored solutions for every need—whether scaling an enterprise or celebrating a milestone."),
		WorkStepsTitle:         models.Str("How We Work"),
		WorkStepsSubtitle:      models.Str("A seamless process designed to save you time and ensure top-quality results"),
		BranchesBadge:          models.Str("Lokasi Cabang"),
		BranchesTitle:          models.Str("Temukan Cabang Terdekat Kami"),
		BranchesSubtitle:       models.Str("Kunjungi gerai fisik kami untuk berkonsultasi langsung atau mengambil pesanan Anda."),
		ClientsTitle:           models.Str("Klien Kami"),
		ClientsSubtitle:        models.Str("Telah dipercaya oleh berbagai perusahaan dan institusi di Indonesia untuk solusi bisnis berkualitas."),
		CreatedAt:              now(),
		UpdatedAt:              now(),
	}).FirstOrCreate(&site)
	if res.Error != nil {
		return nil, nil, nil, nil, nil, nil, res.Error
	}
	if res.RowsAffected > 0 {
		_ = database.ResetSequence(db, "site_configs")
	}

	hero := models.HeroSection{}
	res = db.Where(models.HeroSection{ID: 1}).Attrs(models.HeroSection{
		Badge:            models.Str("Solusi Terpercaya"),
		Title:            "Wujudkan Solusi Terbaik Untuk Bisnis Anda",
		Subtitle:         models.Str("Temukan solusi profesional dan terintegrasi berkualitas tinggi untuk mendukung pertumbuhan bisnis Anda."),
		PrimaryBtnText:   models.Str("Mulai Proyek"),
		PrimaryBtnURL:    models.Str("whatsapp"),
		SecondaryBtnText: models.Str("Layanan Kami"),
		SecondaryBtnURL:  models.Str("#services"),
		CreatedAt:        now(),
		UpdatedAt:        now(),
	}).FirstOrCreate(&hero)
	if res.Error != nil {
		return nil, nil, nil, nil, nil, nil, res.Error
	}
	if res.RowsAffected > 0 {
		_ = database.ResetSequence(db, "hero_sections")
	}

	heroBg := models.HeroBackground{}
	res = db.Where(models.HeroBackground{ID: 1}).Attrs(models.HeroBackground{
		Badge:            models.Str("Solusi Digital"),
		Title:            "Wujudkan Solusi Terbaik Untuk Bisnis Anda",
		Subtitle:         models.Str("Temukan solusi profesional dan terintegrasi berkualitas tinggi untuk mendukung pertumbuhan bisnis Anda."),
		PrimaryBtnText:   models.Str("Mulai Proyek"),
		PrimaryBtnURL:    models.Str("whatsapp"),
		SecondaryBtnText: models.Str("Layanan Kami"),
		SecondaryBtnURL:  models.Str("#services"),
		CreatedAt:        now(),
		UpdatedAt:        now(),
	}).FirstOrCreate(&heroBg)
	if res.Error != nil {
		return nil, nil, nil, nil, nil, nil, res.Error
	}
	if res.RowsAffected > 0 {
		_ = database.ResetSequence(db, "hero_backgrounds")
	}

	heroCarousel := models.HeroCarousel{}
	res = db.Where(models.HeroCarousel{ID: 1}).Attrs(models.HeroCarousel{
		Badge:            models.Str("Inovasi Digital"),
		Title:            "Wujudkan Solusi Terbaik Untuk Bisnis Anda",
		Subtitle:         models.Str("Temukan solusi profesional dan terintegrasi berkualitas tinggi untuk mendukung pertumbuhan bisnis Anda."),
		PrimaryBtnText:   models.Str("Mulai Proyek"),
		PrimaryBtnURL:    models.Str("whatsapp"),
		SecondaryBtnText: models.Str("Layanan Kami"),
		SecondaryBtnURL:  models.Str("#services"),
		CarouselImages:   models.StringList{},
		CreatedAt:        now(),
		UpdatedAt:        now(),
	}).FirstOrCreate(&heroCarousel)
	if res.Error != nil {
		return nil, nil, nil, nil, nil, nil, res.Error
	}
	if res.RowsAffected > 0 {
		_ = database.ResetSequence(db, "hero_carousels")
	}

	why := models.WhyChooseUs{}
	res = db.Where(models.WhyChooseUs{ID: 1}).Attrs(models.WhyChooseUs{
		Title:    "Mengapa Memilih Kami?",
		Subtitle: models.Str("Prioritas utama kami adalah memberikan solusi profesional berkualitas tinggi, pengerjaan tepat waktu, dan pelayanan terbaik untuk Anda."),
		Features: models.StringList{
			"Solusi terintegrasi dan berkualitas tinggi",
			"Tim profesional dan berpengalaman",
			"Pengerjaan tepat waktu dengan standar industri",
			"Dukungan teknis responsif dan terpercaya",
		},
		CreatedAt: now(),
		UpdatedAt: now(),
	}).FirstOrCreate(&why)
	if res.Error != nil {
		return nil, nil, nil, nil, nil, nil, res.Error
	}
	if res.RowsAffected > 0 {
		_ = database.ResetSequence(db, "why_choose_us")
	}
	if why.Features == nil {
		why.Features = models.StringList{}
	}

	cta := models.CtaSection{}
	res = db.Where(models.CtaSection{ID: 1}).Attrs(models.CtaSection{
		Title:     "Siap Mewujudkan Proyek Anda?",
		Subtitle:  models.Str("Konsultasikan kebutuhan bisnis Anda bersama tim profesional kami dan temukan solusi terbaik sekarang."),
		BtnText:   models.Str("Mulai Konsultasi"),
		BtnURL:    models.Str("whatsapp"),
		CreatedAt: now(),
		UpdatedAt: now(),
	}).FirstOrCreate(&cta)
	if res.Error != nil {
		return nil, nil, nil, nil, nil, nil, res.Error
	}
	if res.RowsAffected > 0 {
		_ = database.ResetSequence(db, "cta_sections")
	}

	return &site, &hero, &heroBg, &heroCarousel, &why, &cta, nil
}

// ListDefaults seeds default rows for list tables that are empty.
func ListDefaults(db *gorm.DB) error {
	var count int64

	if err := db.Model(&models.Statistic{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		rows := []models.Statistic{
			{Value: "500+", Label: "Klien Puas", SortOrder: 1},
			{Value: "1,200+", Label: "Proyek Selesai", SortOrder: 2},
			{Value: "5 Tahun", Label: "Pengalaman", SortOrder: 3},
			{Value: "4.9/5", Label: "Rating Klien", SortOrder: 4},
		}
		for i := range rows {
			rows[i].CreatedAt, rows[i].UpdatedAt = now(), now()
		}
		if err := db.Create(&rows).Error; err != nil {
			return err
		}
	}

	if err := db.Model(&models.Service{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		rows := []models.Service{
			{Title: "Website & App Development", Description: "Solusi pembuatan website profil perusahaan, aplikasi web, dan landing page modern berperforma tinggi.", Icon: "globe", SortOrder: 1},
			{Title: "UI/UX & Branding", Description: "Desain identitas visual, antarmuka pengguna elegan, dan pengalaman digital yang mengesankan.", Icon: "layout", SortOrder: 2},
			{Title: "Custom Dashboards & ERP", Description: "Pengembangan sistem manajemen data, CRM, dan sistem otomasi bisnis terintegrasi.", Icon: "cpu", SortOrder: 3},
		}
		for i := range rows {
			rows[i].CreatedAt, rows[i].UpdatedAt = now(), now()
		}
		if err := db.Create(&rows).Error; err != nil {
			return err
		}
	}

	if err := db.Model(&models.OrderStep{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		rows := []models.OrderStep{
			{StepNumber: 1, Title: "Konsultasi", Description: "Hubungi kami via WhatsApp atau form kontak untuk mendiskusikan kebutuhan dan spesifikasi proyek Anda.", Icon: models.Str("message-square"), SortOrder: 1},
			{StepNumber: 2, Title: "Perencanaan & Desain", Description: "Kirim brief kebutuhan atau diskusikan konsep bersama tim ahli kami untuk rencana terbaik.", Icon: models.Str("edit"), SortOrder: 2},
			{StepNumber: 3, Title: "Pengerjaan Proyek", Description: "Proses pengerjaan dilakukan secara profesional dengan standar mutu tinggi dan pengawasan berkala.", Icon: models.Str("briefcase"), SortOrder: 3},
			{StepNumber: 4, Title: "Selesai & Serah Terima", Description: "Hasil proyek siap diserahkan dan diimplementasikan sesuai kebutuhan bisnis Anda.", Icon: models.Str("check"), SortOrder: 4},
		}
		for i := range rows {
			rows[i].CreatedAt, rows[i].UpdatedAt = now(), now()
		}
		if err := db.Create(&rows).Error; err != nil {
			return err
		}
	}

	if err := db.Model(&models.Testimonial{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		rows := []models.Testimonial{
			{ClientName: "Rian Diantono", ClientRole: models.Str("Pemilik Kedai Kopi"), Stars: 5, Content: "Sangat puas dengan layanan dan hasil proyek. Pengerjaannya sangat cepat, rapi, dan timnya sangat komunikatif!", SortOrder: 1},
			{ClientName: "Sari Dewi", ClientRole: models.Str("Kreatif Agensi"), Stars: 5, Content: "Desain dan fungsionalitas sistem melebihi ekspektasi kami. Sangat direkomendasikan untuk solusi profesional.", SortOrder: 2},
			{ClientName: "Yusuf Wibowo", ClientRole: models.Str("Project Manager"), Stars: 5, Content: "Dukungan teknis responsif dan hasil akhir stabil. Pelayanan sangat ramah dan profesional.", SortOrder: 3},
		}
		for i := range rows {
			rows[i].CreatedAt, rows[i].UpdatedAt = now(), now()
		}
		if err := db.Create(&rows).Error; err != nil {
			return err
		}
	}

	if err := db.Model(&models.LandingSection{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		rows := []models.LandingSection{
			{SectionKey: "hero", Title: "Hero (Judul Utama & Banner)", IsActive: true, SortOrder: 1},
			{SectionKey: "statistics", Title: "Statistik Bar (Pencapaian)", IsActive: true, SortOrder: 2},
			{SectionKey: "services", Title: "Layanan Kami (Services)", IsActive: true, SortOrder: 3},
			{SectionKey: "benefits", Title: "Keunggulan Kami (Why Choose Us)", IsActive: true, SortOrder: 4},
			{SectionKey: "portfolio", Title: "Galeri Portofolio", IsActive: true, SortOrder: 5},
			{SectionKey: "timeline", Title: "Alur Kerja (Work Steps)", IsActive: true, SortOrder: 6},
			{SectionKey: "testimonials", Title: "Testimoni Klien & Mitra", IsActive: true, SortOrder: 7},
			{SectionKey: "cta", Title: "Banner Hubungi Kami (CTA)", IsActive: true, SortOrder: 8},
		}
		for i := range rows {
			rows[i].CreatedAt, rows[i].UpdatedAt = now(), now()
		}
		if err := db.Create(&rows).Error; err != nil {
			return err
		}
	}

	if err := db.Model(&models.AboutItem{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		rows := []models.AboutItem{
			{Title: "Visionary Architecture", Description: models.Str("Scalable, future-proof engineering designed to grow alongside your ambitions."), SortOrder: 1},
			{Title: "Human-Centric UX", Description: models.Str("Intuitive workflows and high-converting aesthetics that resonate with modern users."), SortOrder: 2},
			{Title: "Agile Precision", Description: models.Str("Iterative execution backed by weekly feedback loops and dedicated support."), SortOrder: 3},
			{Title: "Enterprise Security", Description: models.Str("Built from day one with modern compliance, encryption, and performance standards."), SortOrder: 4},
		}
		for i := range rows {
			rows[i].CreatedAt, rows[i].UpdatedAt = now(), now()
		}
		if err := db.Create(&rows).Error; err != nil {
			return err
		}
	}

	if err := db.Model(&models.ServicePremium{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		rows := []models.ServicePremium{
			{Title: "Landing Pages", Description: models.Str("High-converting, responsive landing pages designed to capture leads and elevate brand authority."), ButtonText: "Start Project", ButtonURL: "whatsapp", SortOrder: 1},
			{Title: "Company Profiles", Description: models.Str("Professional web presence that communicates credibility, vision, and core corporate value."), ButtonText: "Start Project", ButtonURL: "whatsapp", SortOrder: 2},
			{Title: "Custom Dashboards", Description: models.Str("Data management, ERP, and CRM solutions tailored to streamline your business operations."), ButtonText: "Start Project", ButtonURL: "whatsapp", SortOrder: 3},
			{Title: "Wedding Templates", Description: models.Str("Aesthetic, interactive digital invitations to make your special day unforgettable."), ButtonText: "Start Project", ButtonURL: "whatsapp", SortOrder: 4},
		}
		for i := range rows {
			rows[i].CreatedAt, rows[i].UpdatedAt = now(), now()
		}
		if err := db.Create(&rows).Error; err != nil {
			return err
		}
	}

	if err := db.Model(&models.WorkStep{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		rows := []models.WorkStep{
			{Title: "Consultation", Description: models.Str("Share your vision, requirements, and business goals with our expert team."), SortOrder: 1},
			{Title: "Design & Planning", Description: models.Str("We map out the UX flow and visual architecture tailored perfectly to your brand."), SortOrder: 2},
			{Title: "Launch & Support", Description: models.Str("Seamless deployment and ongoing maintenance to keep your product running perfectly."), SortOrder: 3},
			{Title: "Fast Delivery", Description: models.Str("Rapid development with regular updates and continuous feedback loops."), SortOrder: 4},
		}
		for i := range rows {
			rows[i].CreatedAt, rows[i].UpdatedAt = now(), now()
		}
		if err := db.Create(&rows).Error; err != nil {
			return err
		}
	}

	if err := db.Model(&models.Branch{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		b := models.Branch{
			Name:      "Kantor Pusat Jakarta",
			Address:   "Jl. Jenderal Sudirman No. 10, Jakarta Pusat, DKI Jakarta",
			Latitude:  "-6.208763",
			Longitude: "106.845599",
			Phone:     models.Str("628123456789"),
			IsActive:  true,
			SortOrder: 1,
			CreatedAt: now(),
			UpdatedAt: now(),
		}
		if err := db.Create(&b).Error; err != nil {
			return err
		}
	}

	if err := db.Model(&models.Client{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		cl := models.Client{
			Name:       "Maju Jaya Company",
			Logo:       "clients/default.png",
			WebsiteURL: models.Str("https://google.com"),
			IsActive:   true,
			SortOrder:  1,
			CreatedAt:  now(),
			UpdatedAt:  now(),
		}
		if err := db.Create(&cl).Error; err != nil {
			return err
		}
	}

	return nil
}
