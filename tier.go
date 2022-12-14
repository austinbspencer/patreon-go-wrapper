package patreon

// TierFields is all fields in the Tier Attributes struct
var TierFields = getObjectFields(Tier{}.Attributes)

// Tier is a membership level on a campaign, which can have benefits attached to it.
type Tier struct {
	Type          string         `json:"type"`
	ID            string         `json:"id"`
	Attributes    TierAttributes `json:"attributes"`
	Relationships struct {
		Benefits  *BenefitsRelationship `json:"benefits,omitempty"`
		Campaign  *CampaignRelationship `json:"campaign,omitempty"`
		TierImage *MediaRelationship    `json:"tier_image,omitempty"`
	} `json:"relationships"`
}

// TierAttributes is the attributes struct for Tier
type TierAttributes struct {
	AmountCents      int         `json:"amount_cents"`
	CreatedAt        NullTime    `json:"created_at"`
	Description      string      `json:"description"`
	DiscordRoleIds   interface{} `json:"discord_role_ids"`
	EditedAt         NullTime    `json:"edited_at"`
	ImageURL         string      `json:"image_url"`
	PatronCount      int         `json:"patron_count"`
	PostCount        int         `json:"post_count"`
	Published        bool        `json:"published"`
	PublishedAt      NullTime    `json:"published_at"`
	Remaining        int         `json:"remaining"`
	RequiresShipping bool        `json:"requires_shipping"`
	Title            string      `json:"title"`
	UnpublishedAt    NullTime    `json:"unpublished_at"`
	URL              string      `json:"url"`
	UserLimit        int         `json:"user_limit"`
}
