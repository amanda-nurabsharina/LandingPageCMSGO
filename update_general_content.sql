-- Update site_configs to General CMS defaults
UPDATE site_configs SET
  site_name = 'FourPlusOne',
  contact_title = 'Kirimkan Pesan atau Konsultasi Gratis',
  contact_subtitle = 'Punya pertanyaan seputar layanan kami, atau ingin mendiskusikan kebutuhan proyek Anda? Isi formulir, tim kami akan segera menghubungi Anda.',
  services_title = 'Solusi Profesional & Layanan Terintegrasi',
  services_subtitle = 'Kami siap menghadirkan berbagai solusi terbaik untuk kebutuhan branding, pengembangan bisnis, dan transformasi digital Anda.',
  order_steps_title = 'Alur Kerja Mudah & Terstruktur',
  order_steps_subtitle = 'Cukup selesaikan 4 langkah mudah berikut untuk mewujudkan proyek dan kebutuhan bisnis Anda.',
  testimonials_title = 'Apa Kata Klien & Mitra Kami',
  testimonials_subtitle = 'Kepercayaan dan kepuasan klien adalah komitmen utama kami. Simak penilaian langsung mereka.',
  news_title = 'Berita & Informasi Terkini',
  news_subtitle = 'Ikuti perkembangan terbaru mengenai layanan, inovasi, dan tips seputar industri kami.',
  activities_title = 'Aktivitas & Dokumentasi',
  activities_subtitle = 'Dokumentasi portofolio kerja, aktivitas tim, serta event penting yang kami hadiri.',
  footer_description = 'Menyediakan layanan profesional dan solusi digital berkualitas tinggi untuk mendukung pertumbuhan bisnis Anda.'
WHERE id = 1;

-- Update hero_sections
UPDATE hero_sections SET
  badge = 'Solusi Terpercaya',
  title = 'Wujudkan Solusi Terbaik Untuk Bisnis Anda',
  subtitle = 'Temukan solusi profesional dan terintegrasi berkualitas tinggi untuk mendukung pertumbuhan bisnis Anda.',
  primary_btn_text = 'Mulai Proyek',
  secondary_btn_text = 'Layanan Kami'
WHERE id = 1;

-- Update hero_backgrounds
UPDATE hero_backgrounds SET
  badge = 'Solusi Digital',
  title = 'Wujudkan Solusi Terbaik Untuk Bisnis Anda',
  subtitle = 'Temukan solusi profesional dan terintegrasi berkualitas tinggi untuk mendukung pertumbuhan bisnis Anda.',
  primary_btn_text = 'Mulai Proyek',
  secondary_btn_text = 'Layanan Kami'
WHERE id = 1;

-- Update why_choose_us
UPDATE why_choose_us SET
  title = 'Mengapa Memilih Kami?',
  subtitle = 'Prioritas utama kami adalah memberikan hasil dengan kualitas premium, pengerjaan cepat & tepat, serta pelayanan terbaik untuk Anda.',
  features = '["Kualitas hasil berstandar tinggi","Tim ahli & berpengalaman","Pengerjaan cepat & tepat waktu","Solusi terpercaya & kompetitif"]'
WHERE id = 1;

-- Update services
UPDATE services SET
  title = 'Website & Web Application',
  description = 'Pengembangan website perusahaan, web portal responsif, dan landing page modern berperforma tinggi.',
  icon = 'globe'
WHERE id = 1;

UPDATE services SET
  title = 'UI/UX & Branding Design',
  description = 'Perancangan identitas visual, antarmuka pengguna elegan, dan pengalaman interaktif yang intuitif.',
  icon = 'layout'
WHERE id = 2;

UPDATE services SET
  title = 'Custom Dashboards & ERP',
  description = 'Solusi manajemen data internal, CRM, dan sistem otomasi alur kerja bisnis yang efisien.',
  icon = 'cpu'
WHERE id = 3;

UPDATE services SET
  title = 'Mobile App Development',
  description = 'Aplikasi mobile iOS & Android yang tangguh, aman, dan dirancang untuk pengalaman pengguna maksimal.',
  icon = 'smartphone'
WHERE id = 4;

UPDATE services SET
  title = 'Cloud & DevOps Solutions',
  description = 'Infrastruktur cloud andal, arsitektur microservices, konfigurasi server aman, dan automasi deployment.',
  icon = 'cloud'
WHERE id = 5;

UPDATE services SET
  title = 'Maintenance & Tech Support',
  description = 'Dukungan pemeliharaan sistem berkala, audit performa, perbaikan keamanan, dan asistensi teknis berkelanjutan.',
  icon = 'headphones'
WHERE id = 6;

-- Update order_steps
UPDATE order_steps SET
  title = 'Konsultasi',
  description = 'Hubungi kami via WhatsApp atau form kontak untuk mendiskusikan kebutuhan dan spesifikasi proyek Anda.',
  icon = 'message-square'
WHERE id = 1;

UPDATE order_steps SET
  title = 'Perencanaan & Desain',
  description = 'Kirim brief kebutuhan atau diskusikan konsep bersama tim ahli kami untuk rencana terbaik.',
  icon = 'edit'
WHERE id = 2;

UPDATE order_steps SET
  title = 'Pengerjaan Proyek',
  description = 'Proses pengerjaan dilakukan secara profesional dengan standar mutu tinggi dan pengawasan berkala.',
  icon = 'briefcase'
WHERE id = 3;

UPDATE order_steps SET
  title = 'Selesai & Serah Terima',
  description = 'Hasil proyek siap diserahkan dan diimplementasikan sesuai kebutuhan bisnis Anda.',
  icon = 'check'
WHERE id = 4;

-- Update testimonials
UPDATE testimonials SET
  client_name = 'Rian Diantono',
  client_role = 'Founder Start-Up',
  content = 'Sangat puas dengan kualitas pengerjaan dan hasil proyek dari tim ini. Komunikasi responsif, desain modern, dan pengerjaannya sangat cepat!'
WHERE id = 1;

UPDATE testimonials SET
  client_name = 'Sari Dewi',
  client_role = 'Creative Director',
  content = 'Website dan landing page yang dibuat memiliki performa kencang dan desain UI yang elegan. Sangat merekomendasikan layanan profesional ini.'
WHERE id = 2;

UPDATE testimonials SET
  client_name = 'Yusuf Wibowo',
  client_role = 'Operations Lead',
  content = 'Sistem dashboard custom yang dikembangkan sangat membantu menyederhanakan alur kerja tim kami. Pelayanan dan dukungan teknisnya sangat memuaskan.'
WHERE id = 3;

-- Update statistics
UPDATE statistics SET
  value = '1,200+',
  label = 'Proyek Selesai'
WHERE id = 2;

-- Update landing_sections
UPDATE landing_sections SET title = 'Layanan Kami (Services)' WHERE section_key = 'services';
UPDATE landing_sections SET title = 'Alur Kerja (Work Steps)' WHERE section_key = 'timeline';
UPDATE landing_sections SET title = 'Testimoni Klien & Mitra' WHERE section_key = 'testimonials';

