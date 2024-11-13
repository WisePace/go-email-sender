
# Pace Email Sender

Pace Email Sender is a Go-based application designed to send emails to a list of valid email addresses.
The application reads a list of emails from a text file (db.txt), validates them, and then sends a predefined email message to those recipients using SMTP (Gmail's SMTP server).

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


**Update the main.go file with your Gmail account details and app-specific password:**

```
    smtpHost := "smtp.gmail.com"
    smtpPort := "587"
    sender := "your-email@gmail.com"   // Your Gmail account
    password := "your-app-password"    // App-specific password for Gmail
```

## Running the Application

**To run the application, simply use the following command:**

```
go run main.go

```

The application will:

<ol><li>Read and validate email addresses from db.txt. </li>
    <li>Read the email body from letter.txt.</li>
    <li>Send the email to all valid email addresses concurrently (with a limit on the number of concurrent goroutines).</li>
    </ol>