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

namespace prevention_productivity.Pages.SchoolReports
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
            return Page();
        }

        [BindProperty]
        public SchoolReport SchoolReport { get; set; }

        // To protect from overposting attacks, see https://aka.ms/RazorPagesCRUD
        public async Task<IActionResult> OnPostAsync()
        {
            SchoolReport.TeamMember.Id = UserManager.GetUserId(User);
            SchoolReport.Status = ApprovalStatus.Pending;

            var isAuthorized = await AuthorizationService.AuthorizeAsync(
                                                        User, SchoolReport,
                                                        AuthOperations.Create);
            if (!isAuthorized.Succeeded)
            {
                return Forbid();
            }
            _context.SchoolReport.Add(SchoolReport);
            await _context.SaveChangesAsync();

            return RedirectToPage("./Index");
        }
    }
}
