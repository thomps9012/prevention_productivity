#nullable disable
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Microsoft.EntityFrameworkCore;
using prevention_productivity.Authorization;
using prevention_productivity.Data;
using prevention_productivity.Models;
using prevention_productivity.Pages.ProductivityLogs;

namespace prevention_productivity.Pages.Contacts
{
    public class IndexModel : DI_BasePageModel
    {
        private readonly ApplicationDbContext _context;

        public IndexModel(ApplicationDbContext context,
             IAuthorizationService authorizationService,
            UserManager<ApplicationUser> userManager)
            : base(context, authorizationService, userManager)
        {
            _context = context;
        }

        public IList<Contact> Contact { get;set; }
        public string TypeSearch { get; set; }

        public async Task OnGetAsync(string typeSearch)
        {
            var contacts = from c in _context.Contact
                           select c;

            var isAuthorized = User.IsInRole(Constants.AdminRole);
            
            if(!isAuthorized)
            {
                contacts = contacts.Where(c => 
                c.Type == ContactType.Student 
                || c.Type == ContactType.Parent);
            }
            switch (typeSearch)
            {
                case "Student":
                    contacts = contacts.Where(c => c.Type == ContactType.Student);
                    break;
                case "Parent":
                    contacts = contacts.Where(c => c.Type == ContactType.Parent);
                    break;
                case "Teacher":
                    contacts = contacts.Where(c => c.Type == ContactType.Teacher);
                    break;
                case "NonProfit":
                    contacts = contacts.Where(c => c.Type == ContactType.NonProfit);
                    break;
                case "Public":
                    contacts = contacts.Where(c => c.Type == ContactType.Public);
                    break;
                case "Private":
                    contacts = contacts.Where(c => c.Type == ContactType.Private);
                    break;
                case "Other":
                    contacts = contacts.Where(c => c.Type == ContactType.Other);
                    break;
                default:
                    break;
            }
            Contact = await contacts.ToListAsync();
        }
    }
}
