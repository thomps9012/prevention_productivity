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
    public class DeleteModel : DI_BasePageModel
    {
        private readonly ApplicationDbContext _context;

        public DeleteModel(ApplicationDbContext context,
            IAuthorizationService authorizationService,
            UserManager<ApplicationUser> userManager)
            : base(context, authorizationService, userManager)
        {
            _context = context;
        }

        [BindProperty]
        public Contact Contact { get; set; }

        public async Task<IActionResult> OnGetAsync(int? id)
        {
            Contact? _contact = await _context.Contact.FirstOrDefaultAsync(m => m.ContactId == id);
            if (_contact == null)
            {
                return NotFound();
            }
            Contact = _contact;
            var isAuthorized = await AuthorizationService.AuthorizeAsync(User, Contact,
                                                                AuthOperations.Delete);
            if (!isAuthorized.Succeeded)
            {
                return Forbid();
            }
            return Page();
        }

        public async Task<IActionResult> OnPostAsync(int? id)
        {
            var contact = await _context.Contact.FindAsync(id);
            if (contact == null)
            {
                return NotFound();
            }

            var isAuthorized = await AuthorizationService.AuthorizeAsync(
                User, Contact, AuthOperations.Delete);

            if (!isAuthorized.Succeeded)
            {
                return Forbid();
            }
                _context.Contact.Remove(contact);
                await _context.SaveChangesAsync();

            return RedirectToPage("./Index");
        }
    }
}
