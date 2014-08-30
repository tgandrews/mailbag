package subscriber

import "time"

// A Subscriber of the list
type Subscriber struct {
	// The time they subscribed
	SubscribeTime time.Time

	// The user agent of the browser requested
	UserAgent string

	// The HTTP referer
	Referer string

	// The IP address the email was registered with
	IPAddress string

	// The email address
	EmailAddress string
}

// ToCSV turns the subscriber to the format for CSV
func (s Subscriber) ToCSV() []string {
	return []string{
		s.EmailAddress,
		s.SubscribeTime.String(),
		s.UserAgent,
		s.IPAddress}
}
