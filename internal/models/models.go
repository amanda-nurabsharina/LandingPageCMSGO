package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// StringList is a []string persisted as a JSON array (json in PostgreSQL).
type StringList []string

func (s StringList) Value() (driver.Value, error) {
	if s == nil {
		s = StringList{}
	}
	b, err := json.Marshal([]string(s))
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

func (s *StringList) Scan(src any) error {
	var raw []byte
	switch v := src.(type) {
	case nil:
		*s = StringList{}
		return nil
	case []byte:
		raw = v
	case string:
		raw = []byte(v)
	default:
		return fmt.Errorf("StringList: unsupported type %T", src)
	}
	if len(raw) == 0 {
		*s = StringList{}
		return nil
	}
	var out []string
	if err := json.Unmarshal(raw, &out); err != nil {
		return err
	}
	if out == nil {
		out = []string{}
	}
	*s = out
	return nil
}

func (StringList) GormDataType() string { return "json" }

// ---------------------------------------------------------------------------
// Models (Branch General)
// ---------------------------------------------------------------------------

type SiteConfig struct {
	ID                     uint       `gorm:"primaryKey" json:"id"`
	SiteName               string     `gorm:"column:site_name;type:varchar(255);not null;default:'FourPlusOne'" json:"site_name"`
	Logo                   *string    `gorm:"column:logo;type:varchar(255)" json:"logo"`
	WhatsappNumber         string     `gorm:"column:whatsapp_number;type:varchar(255);not null;default:'628123456789'" json:"whatsapp_number"`
	Email                  *string    `gorm:"column:email;type:varchar(255)" json:"email"`
	Address                *string    `gorm:"column:address;type:text" json:"address"`
	FacebookURL            *string    `gorm:"column:facebook_url;type:varchar(255)" json:"facebook_url"`
	InstagramURL           *string    `gorm:"column:instagram_url;type:varchar(255)" json:"instagram_url"`
	TwitterURL             *string    `gorm:"column:twitter_url;type:varchar(255)" json:"twitter_url"`
	PrimaryColor           *string    `gorm:"column:primary_color;type:varchar(255)" json:"primary_color"`
	SecondaryColor         *string    `gorm:"column:secondary_color;type:varchar(255)" json:"secondary_color"`
	AccentColor            *string    `gorm:"column:accent_color;type:varchar(255)" json:"accent_color"`
	BackgroundColor        *string    `gorm:"column:background_color;type:varchar(255)" json:"background_color"`
	ContactTitle           *string    `gorm:"column:contact_title;type:varchar(255)" json:"contact_title"`
	ContactSubtitle        *string    `gorm:"column:contact_subtitle;type:text" json:"contact_subtitle"`
	NewsTitle              *string    `gorm:"column:news_title;type:varchar(255)" json:"news_title"`
	NewsSubtitle           *string    `gorm:"column:news_subtitle;type:text" json:"news_subtitle"`
	ActivitiesTitle        *string    `gorm:"column:activities_title;type:varchar(255)" json:"activities_title"`
	ActivitiesSubtitle     *string    `gorm:"column:activities_subtitle;type:text" json:"activities_subtitle"`
	ServicesTitle          *string    `gorm:"column:services_title;type:varchar(255)" json:"services_title"`
	ServicesSubtitle       *string    `gorm:"column:services_subtitle;type:text" json:"services_subtitle"`
	PortfolioTitle         *string    `gorm:"column:portfolio_title;type:varchar(255)" json:"portfolio_title"`
	PortfolioSubtitle      *string    `gorm:"column:portfolio_subtitle;type:text" json:"portfolio_subtitle"`
	OrderStepsTitle        *string    `gorm:"column:order_steps_title;type:varchar(255)" json:"order_steps_title"`
	OrderStepsSubtitle     *string    `gorm:"column:order_steps_subtitle;type:text" json:"order_steps_subtitle"`
	TestimonialsTitle      *string    `gorm:"column:testimonials_title;type:varchar(255)" json:"testimonials_title"`
	TestimonialsSubtitle   *string    `gorm:"column:testimonials_subtitle;type:text" json:"testimonials_subtitle"`
	FooterDescription      *string    `gorm:"column:footer_description;type:text" json:"footer_description"`
	AboutTitle             *string    `gorm:"column:about_title;type:varchar(255)" json:"about_title"`
	AboutSubtitle          *string    `gorm:"column:about_subtitle;type:text" json:"about_subtitle"`
	ServicePremiumTitle    *string    `gorm:"column:service_premium_title;type:varchar(255)" json:"service_premium_title"`
	ServicePremiumSubtitle *string    `gorm:"column:service_premium_subtitle;type:text" json:"service_premium_subtitle"`
	WorkStepsTitle         *string    `gorm:"column:work_steps_title;type:varchar(255)" json:"work_steps_title"`
	WorkStepsSubtitle      *string    `gorm:"column:work_steps_subtitle;type:text" json:"work_steps_subtitle"`
	BranchesBadge          *string    `gorm:"column:branches_badge;type:varchar(255)" json:"branches_badge"`
	BranchesTitle          *string    `gorm:"column:branches_title;type:varchar(255)" json:"branches_title"`
	BranchesSubtitle       *string    `gorm:"column:branches_subtitle;type:text" json:"branches_subtitle"`
	ClientsTitle           *string    `gorm:"column:clients_title;type:varchar(255)" json:"clients_title"`
	ClientsSubtitle        *string    `gorm:"column:clients_subtitle;type:text" json:"clients_subtitle"`
	CreatedAt              *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt              *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (SiteConfig) TableName() string { return "site_configs" }

type HeroSection struct {
	ID               uint       `gorm:"primaryKey" json:"id"`
	Badge            *string    `gorm:"column:badge;type:varchar(255)" json:"badge"`
	Title            string     `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Subtitle         *string    `gorm:"column:subtitle;type:text" json:"subtitle"`
	PrimaryBtnText   *string    `gorm:"column:primary_btn_text;type:varchar(255)" json:"primary_btn_text"`
	PrimaryBtnURL    *string    `gorm:"column:primary_btn_url;type:varchar(255)" json:"primary_btn_url"`
	SecondaryBtnText *string    `gorm:"column:secondary_btn_text;type:varchar(255)" json:"secondary_btn_text"`
	SecondaryBtnURL  *string    `gorm:"column:secondary_btn_url;type:varchar(255)" json:"secondary_btn_url"`
	ImagePath        *string    `gorm:"column:image_path;type:varchar(255)" json:"image_path"`
	CreatedAt        *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (HeroSection) TableName() string { return "hero_sections" }

type HeroBackground struct {
	ID               uint       `gorm:"primaryKey" json:"id"`
	Badge            *string    `gorm:"column:badge;type:varchar(255)" json:"badge"`
	Title            string     `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Subtitle         *string    `gorm:"column:subtitle;type:text" json:"subtitle"`
	PrimaryBtnText   *string    `gorm:"column:primary_btn_text;type:varchar(255)" json:"primary_btn_text"`
	PrimaryBtnURL    *string    `gorm:"column:primary_btn_url;type:varchar(255)" json:"primary_btn_url"`
	SecondaryBtnText *string    `gorm:"column:secondary_btn_text;type:varchar(255)" json:"secondary_btn_text"`
	SecondaryBtnURL  *string    `gorm:"column:secondary_btn_url;type:varchar(255)" json:"secondary_btn_url"`
	ImagePath        *string    `gorm:"column:image_path;type:varchar(255)" json:"image_path"`
	CreatedAt        *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (HeroBackground) TableName() string { return "hero_backgrounds" }

type HeroCarousel struct {
	ID               uint       `gorm:"primaryKey" json:"id"`
	Badge            *string    `gorm:"column:badge;type:varchar(255)" json:"badge"`
	Title            string     `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Subtitle         *string    `gorm:"column:subtitle;type:text" json:"subtitle"`
	PrimaryBtnText   *string    `gorm:"column:primary_btn_text;type:varchar(255)" json:"primary_btn_text"`
	PrimaryBtnURL    *string    `gorm:"column:primary_btn_url;type:varchar(255)" json:"primary_btn_url"`
	SecondaryBtnText *string    `gorm:"column:secondary_btn_text;type:varchar(255)" json:"secondary_btn_text"`
	SecondaryBtnURL  *string    `gorm:"column:secondary_btn_url;type:varchar(255)" json:"secondary_btn_url"`
	CarouselImages   StringList `gorm:"column:carousel_images;type:json" json:"carousel_images"`
	CreatedAt        *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (HeroCarousel) TableName() string { return "hero_carousels" }

type WhyChooseUs struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Title     string     `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Subtitle  *string    `gorm:"column:subtitle;type:text" json:"subtitle"`
	Features  StringList `gorm:"column:features;type:json;not null" json:"features"`
	ImagePath *string    `gorm:"column:image_path;type:varchar(255)" json:"image_path"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (WhyChooseUs) TableName() string { return "why_choose_us" }

type CtaSection struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Title     string     `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Subtitle  *string    `gorm:"column:subtitle;type:text" json:"subtitle"`
	BtnText   *string    `gorm:"column:btn_text;type:varchar(255)" json:"btn_text"`
	BtnURL    *string    `gorm:"column:btn_url;type:varchar(255)" json:"btn_url"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (CtaSection) TableName() string { return "cta_sections" }

type Statistic struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Value     string     `gorm:"column:value;type:varchar(255);not null" json:"value"`
	Label     string     `gorm:"column:label;type:varchar(255);not null" json:"label"`
	SortOrder int        `gorm:"column:sort_order;not null;default:0" json:"sort_order"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Statistic) TableName() string { return "statistics" }

type Service struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Description string     `gorm:"column:description;type:text;not null" json:"description"`
	Icon        string     `gorm:"column:icon;type:varchar(255);not null" json:"icon"`
	SortOrder   int        `gorm:"column:sort_order;not null;default:0" json:"sort_order"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Service) TableName() string { return "services" }

type OrderStep struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	StepNumber  int        `gorm:"column:step_number;not null" json:"step_number"`
	Title       string     `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Description string     `gorm:"column:description;type:text;not null" json:"description"`
	Icon        *string    `gorm:"column:icon;type:varchar(255)" json:"icon"`
	SortOrder   int        `gorm:"column:sort_order;not null;default:0" json:"sort_order"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (OrderStep) TableName() string { return "order_steps" }

type Portfolio struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Title     string     `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Category  string     `gorm:"column:category;type:varchar(255);not null" json:"category"`
	ImagePath string     `gorm:"column:image_path;type:varchar(255);not null" json:"image_path"`
	SortOrder int        `gorm:"column:sort_order;not null;default:0" json:"sort_order"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Portfolio) TableName() string { return "portfolios" }

type Testimonial struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	ClientName string     `gorm:"column:client_name;type:varchar(255);not null" json:"client_name"`
	ClientRole *string    `gorm:"column:client_role;type:varchar(255)" json:"client_role"`
	Stars      int        `gorm:"column:stars;not null;default:5" json:"stars"`
	Content    string     `gorm:"column:content;type:text;not null" json:"content"`
	SortOrder  int        `gorm:"column:sort_order;not null;default:0" json:"sort_order"`
	CreatedAt  *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Testimonial) TableName() string { return "testimonials" }

type LandingSection struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	SectionKey string     `gorm:"column:section_key;type:varchar(255);not null;unique" json:"section_key"`
	Title      string     `gorm:"column:title;type:varchar(255);not null" json:"title"`
	IsActive   bool       `gorm:"column:is_active;not null;default:true" json:"is_active"`
	SortOrder  int        `gorm:"column:sort_order;not null;default:0" json:"sort_order"`
	CreatedAt  *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (LandingSection) TableName() string { return "landing_sections" }

type AboutItem struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Description *string    `gorm:"column:description;type:text" json:"description"`
	ImagePath   *string    `gorm:"column:image_path;type:varchar(255)" json:"image_path"`
	SortOrder   int        `gorm:"column:sort_order;not null;default:0" json:"sort_order"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (AboutItem) TableName() string { return "about_items" }

type ServicePremium struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Description *string    `gorm:"column:description;type:text" json:"description"`
	ImagePath   *string    `gorm:"column:image_path;type:varchar(255)" json:"image_path"`
	ButtonText  string     `gorm:"column:button_text;type:varchar(255);not null;default:'Start Project'" json:"button_text"`
	ButtonURL   string     `gorm:"column:button_url;type:varchar(255);not null;default:'whatsapp'" json:"button_url"`
	SortOrder   int        `gorm:"column:sort_order;not null;default:0" json:"sort_order"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (ServicePremium) TableName() string { return "service_premiums" }

type WorkStep struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Description *string    `gorm:"column:description;type:text" json:"description"`
	ImagePath   *string    `gorm:"column:image_path;type:varchar(255)" json:"image_path"`
	SortOrder   int        `gorm:"column:sort_order;not null;default:0" json:"sort_order"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (WorkStep) TableName() string { return "work_steps" }

type News struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Slug        string     `gorm:"column:slug;type:varchar(255);not null;unique" json:"slug"`
	Thumbnail   *string    `gorm:"column:thumbnail;type:varchar(255)" json:"thumbnail"`
	Description *string    `gorm:"column:description;type:text" json:"description"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (News) TableName() string { return "news" }

type Activity struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Slug        string     `gorm:"column:slug;type:varchar(255);not null;unique" json:"slug"`
	Thumbnail   *string    `gorm:"column:thumbnail;type:varchar(255)" json:"thumbnail"`
	Description *string    `gorm:"column:description;type:text" json:"description"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Activity) TableName() string { return "activities" }

type Branch struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Address   string     `gorm:"column:address;type:text;not null" json:"address"`
	Latitude  string     `gorm:"column:latitude;type:varchar(50);not null" json:"latitude"`
	Longitude string     `gorm:"column:longitude;type:varchar(50);not null" json:"longitude"`
	Phone     *string    `gorm:"column:phone;type:varchar(50)" json:"phone"`
	IsActive  bool       `gorm:"column:is_active;not null;default:true" json:"is_active"`
	SortOrder int        `gorm:"column:sort_order;not null;default:0" json:"sort_order"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Branch) TableName() string { return "branches" }

type Client struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	Name       string     `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Logo       string     `gorm:"column:logo;type:varchar(255);not null" json:"logo"`
	WebsiteURL *string    `gorm:"column:website_url;type:varchar(255)" json:"website_url"`
	IsActive   bool       `gorm:"column:is_active;not null;default:true" json:"is_active"`
	SortOrder  int        `gorm:"column:sort_order;not null;default:0" json:"sort_order"`
	CreatedAt  *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Client) TableName() string { return "clients" }

type Lead struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Phone     string     `gorm:"column:phone;type:varchar(255);not null" json:"phone"`
	Message   string     `gorm:"column:message;type:text;not null" json:"message"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Lead) TableName() string { return "leads" }

type AnalyticsEvent struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	EventType string     `gorm:"column:event_type;type:varchar(50);not null" json:"event_type"`
	PageName  *string    `gorm:"column:page_name;type:varchar(255)" json:"page_name"`
	SessionID string     `gorm:"column:session_id;type:varchar(255);not null" json:"session_id"`
	IPAddress *string    `gorm:"column:ip_address;type:varchar(50)" json:"ip_address"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (AnalyticsEvent) TableName() string { return "analytics_events" }

type User struct {
	ID              uint       `gorm:"primaryKey" json:"id"`
	Name            string     `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Email           string     `gorm:"column:email;type:varchar(255);not null;unique" json:"email"`
	EmailVerifiedAt *time.Time `gorm:"column:email_verified_at" json:"email_verified_at"`
	Password        string     `gorm:"column:password;type:varchar(255);not null" json:"-"`
	RememberToken   *string    `gorm:"column:remember_token;type:varchar(100)" json:"-"`
	CreatedAt       *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (User) TableName() string { return "users" }

// All returns every model, used for auto migration.
func All() []any {
	return []any{
		&User{},
		&SiteConfig{},
		&HeroSection{},
		&HeroBackground{},
		&HeroCarousel{},
		&WhyChooseUs{},
		&CtaSection{},
		&Statistic{},
		&Service{},
		&OrderStep{},
		&Portfolio{},
		&Testimonial{},
		&LandingSection{},
		&AboutItem{},
		&ServicePremium{},
		&WorkStep{},
		&News{},
		&Activity{},
		&Branch{},
		&Client{},
		&Lead{},
		&AnalyticsEvent{},
	}
}

func Str(s string) *string { return &s }
