
# Pace Email Sender

Pace Email Sender is a Go-based application designed to send emails to a list of valid email addresses.
The application reads a list of emails from a text file (db.txt), validates them, and then sends a predefined email message to those recipients using SMTP.

## Features

- **Email Validation:** Validates email addresses using regex.
- **Concurrent Email Sending:** Sends emails concurrently with a controlled concurrency limit to avoid overwhelming the SMTP server.
- **Message Body:** Email content is read from a file (letter.txt), so you can easily customize the message.

## Setup

**Clone this repository:**

```
    git clone https://github.com/WisePace/pace-sender.git
    cd pace-sender
```
**Configure the db.txt file with a list of email addresses. Each email should be on a new line:**


```
    example1@gmail.com
    example2@gmail.com
    example3@gmail.com
```


**Configure the letter.txt file with the message you want to send:**

```
    Hello,

    This is a test email from the Pace Email Sender application.

    Best regards,
    Your Pace Email Sender Team
```


**Update your .env file with your SMTP server details:**

```
    SMTP_HOST=smtp.example.com       # Your SMTP server (e.g., smtp.gmail.com, smtp.mail.yahoo.com)
    SMTP_PORT=587                    # SMTP server port
    SMTP_SENDER=your-email@example.com  # Your email address
    SMTP_PASSWORD=your-email-password   # Your email account password or app-specific password
```

## Running the Application

**To run the application, simply use the following command:**

```
go run main.go

```
