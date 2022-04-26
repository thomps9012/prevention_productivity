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

namespace prevention_productivity.Pages.Events.Summary
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
            //possible change here
        ViewData["EventId"] = new SelectList(_context.Event, "Id", "Name");
            return Page();
        }

        [BindProperty]
        public EventSummary EventSummary { get; set; }

        // To protect from overposting attacks, see https://aka.ms/RazorPagesCRUD
        public async Task<IActionResult> OnPostAsync()
        {
            EventSummary.TeamMemberID = UserManager.GetUserId(User);
            EventSummary.Status = ApprovalStatus.Pending;

            var isAuthorized = await AuthorizationService.AuthorizeAsync(
                User, EventSummary,
                AuthOperations.Create);
            if (!isAuthorized.Succeeded)
            {
                return Forbid();
            }

            _context.EventSummary.Add(EventSummary);
            await _context.SaveChangesAsync();

            return RedirectToPage("./Index");
        }
    }
}
