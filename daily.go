package main

import (
	// "time"
	"os"

	"github.com/s3ththompson/berliner"
	s "github.com/s3ththompson/berliner/sources"
	r "github.com/s3ththompson/berliner/renderers"
	f "github.com/s3ththompson/berliner/filters"
)

func main() {
	b := berliner.New()
	// Add sources
	newYorker := b.Source(s.NewYorker())
	// b.Source(s.DisegnoDaily())
	// Alternatively, read posts from JSON file (see below)
	// b.Source(s.FromJSON("berliner.json"))


	// Add filters to specific sources
	newYorker.Filter(f.Clamp(3)) // Clamp filter allows only n posts

	// Dedupe filter persists list of seen posts to disk
	// b.Filter(f.Dedupe("~/berliner/dedupe"))
	// ReadFor filters posts to those which can be read in 20 minutes at x wpm
	// b.Filter(f.ReadFor(20*time.Minute, 250))
	b.Renderer(r.Terminal())
	// Render HTML to output file; specify template name and css name (found in assets folder)
	b.Renderer(r.HTML("berliner.html", "bootstrap", "bootstrap"))

	// Render to Email using Gmail SMTP server
	emailParams := r.EmailParams{
		SmtpServer: "smtp.gmail.com",
		SmtpPort: 465,
		SmtpUsername: "gklitt@gmail.com",
		SmtpPassword: os.Getenv("SMTP_PASSWORD"),
		FromAddress: "gklitt@gmail.com",
		ToAddress: "gklitt@gmail.com",
	}
	b.Renderer(r.Email(emailParams, "bootstrap", "bootstrap"))
	// Or, write posts to JSON for future use
	// b.Renderer(r.ToJSON("berliner.json"))
	b.Go()
}
