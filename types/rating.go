package types

// Rating identifies the danbooru maturity rating of a post
type Rating string

const (

	// Safe indicates that the post in question is generally lacking sexual content
	// For detailed guidelines, see: https://danbooru.donmai.us/wiki_pages/howto%3Arate
	Safe Rating = "s"

	// Questionable indicates that the post may contain non-genital nudity
	// For detailed guidelines, see: https://danbooru.donmai.us/wiki_pages/howto%3Arate
	Questionable Rating = "q"

	// Explicit indicates that the post may contain full nudity and other mature content
	// For detailed guidelines, see: https://danbooru.donmai.us/wiki_pages/howto%3Arate
	Explicit Rating = "e"
)
