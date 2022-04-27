#nullable disable
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Microsoft.AspNetCore.Mvc.Rendering;
using Microsoft.EntityFrameworkCore;
using prevention_productivity.Authorization;
using prevention_productivity.Data;
using prevention_productivity.Models;
using prevention_productivity.Pages.ProductivityLogs;

namespace prevention_productivity.Pages.Contacts
{
    public class EditModel : DI_BasePageModel
    {
        private readonly ApplicationDbContext _context;

        public EditModel(ApplicationDbContext context,
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
            var isAuthorized = await AuthorizationService.AuthorizeAsync(
               User, Contact, AuthOperations.Update);
            if (!isAuthorized.Succeeded)
            {
            return Forbid();
            }
            if (id == null)
                {
                    return NotFound();
                }

                Contact = await _context.Contact.FirstOrDefaultAsync(m => m.ContactId == id);

                if (Contact == null)
                {
                    return NotFound();
                }
                return Page();
        
            
        }

        // To protect from overposting attacks, enable the specific properties you want to bind to.
        // For more details, see https://aka.ms/RazorPagesCRUD.
        public async Task<IActionResult> OnPostAsync()
        {
            var isAuthorized = await AuthorizationService.AuthorizeAsync(
                User, Contact, AuthOperations.Update);
            if (!isAuthorized.Succeeded)
            {
                return Forbid();
            }
                
            if (!ModelState.IsValid)
            {
                return Page();
            }

            _context.Attach(Contact).State = EntityState.Modified;

            try
            {
                await _context.SaveChangesAsync();
            }
            catch (DbUpdateConcurrencyException)
            {
                if (!ContactExists(Contact.ContactId))
                {
                    return NotFound();
                }
                else
                {
                    throw;
                }
            }

            return RedirectToPage("./Index");
   
        }

        private bool ContactExists(int id)
        {
            return _context.Contact.Any(e => e.ContactId == id);
        }
    }
}
