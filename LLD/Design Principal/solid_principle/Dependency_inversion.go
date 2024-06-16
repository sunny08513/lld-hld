package solidprinciple

//Consider a simple example where a NotificationService sends notifications using an EmailService:

type EmailService struct{}

func (e *EmailService) Send(to string, message string) {
	// Send email
}

type NotificationService struct {
	emailService *EmailService
}

func (n *NotificationService) Notify(to string, message string) {
	n.emailService.Send(to, message)
}

//In this case, the `NotificationService` directly depends on the `EmailService`, making it difficult to switch to another notification method (e.g., SMS) or test the `NotificationService` in isolation. To follow the Dependency Inversion Principle, we can introduce an interface and depend on that instead:
