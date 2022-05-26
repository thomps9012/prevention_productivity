﻿using Microsoft.AspNetCore.Identity.UI.Services;
using Microsoft.Extensions.Options;
using SendGrid;
using SendGrid.Helpers.Mail;

namespace prevention_productivity.Services
{
    public class EmailSender : IEmailSender
    {
        private readonly ILogger _logger;
        
        public EmailSender(IOptions<AuthMessageSenderOptions> optionsAccessor,
            ILogger<EmailSender> logger)
        {
            Options = optionsAccessor.Value;
            _logger = logger;
        }

        public AuthMessageSenderOptions Options { get; }

        public async Task SendEmailAsync(string toEmail, string subject, string message)
        {
            if(string.IsNullOrEmpty(Environment.GetEnvironmentVariable("SendGridAPIKey")))
            {
                throw new Exception("No SendGridKey");
            }
            await Execute(Environment.GetEnvironmentVariable("SendGridAPIKey"), subject, message, toEmail);
        }

        public async Task Execute(string apiKey, string subject, string message, string toEmail)
        {
            var client = new SendGridClient(apiKey);
            var msg = new SendGridMessage()
            {
                From = new EmailAddress("sthompson@norainc.org", "Email Verification"),
                Subject = subject,
                PlainTextContent = message,
                HtmlContent = message
            };
            msg.AddTo(new EmailAddress(toEmail));

            msg.SetClickTracking(false, false);
            var response = await client.SendEmailAsync(msg);
            _logger.LogInformation(response.IsSuccessStatusCode
                ? $"Email to {toEmail} queued successfully!"
                : $"Failure Email to {toEmail}");
        }
    }
}
