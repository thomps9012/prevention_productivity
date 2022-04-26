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
using prevention_productivity.Authorization;
using prevention_productivity.Data;
using prevention_productivity.Models;
using prevention_productivity.Pages.ProductivityLogs;

namespace prevention_productivity.Pages.Events
{
    public class CreateModel : DI_BasePageModel
    {
        private readonly ApplicationDbContext _context;

        public CreateModel(ApplicationDbContext context,
            IAuthorizationService authorizationService,
            UserManager<ApplicationUser> userManager)
            : base(context, authorizationService, userManager)
        {
            _context = context;
        }

        public IActionResult OnGet()
        {
            // possibly need to replace with old solution
        ViewData["GrantProgramId"] = new SelectList(_context.GrantProgram, "Id", "Name");
            TeamMemberList = _context.Users.ToList();
            return Page();
        }

        [BindProperty]
        public Event Event { get; set; }

        public IList<ApplicationUser> TeamMemberList { get; set; }

        // To protect from overposting attacks, see https://aka.ms/RazorPagesCRUD
        public async Task<IActionResult> OnPostAsync()
        {
            Event.EventLead = UserManager.GetUserId(User);
            Event.Status = ApprovalStatus.Pending;


            var isAuthorized = await AuthorizationService.AuthorizeAsync(
                                                        User, Event,
                                                        AuthOperations.Create);
            if (!isAuthorized.Succeeded)
            {
                return Forbid();
            }

            _context.Event.Add(Event);
            await _context.SaveChangesAsync();

            return RedirectToPage("./Index");
        }
    }
}
