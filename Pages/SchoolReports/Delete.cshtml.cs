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

namespace prevention_productivity.Pages.SchoolReports
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
        public SchoolReport SchoolReport { get; set; }

        public async Task<IActionResult> OnGetAsync(int? id)
        {
            SchoolReport? _report = await _context.SchoolReport.FirstOrDefaultAsync(m => m.SchoolReportId == id);
            if (_report == null)
            {
                return NotFound();
            }

            SchoolReport = _report;
            var isAuthorized = await AuthorizationService.AuthorizeAsync(
                                                                   User, SchoolReport,
                                                                   AuthOperations.Delete);
            if (!isAuthorized.Succeeded)
            {
                return Forbid();
            }
            return Page();
        }

        public async Task<IActionResult> OnPostAsync(int? id)
        {
            var report = await _context.SchoolReport.AsNoTracking()
                        .FirstOrDefaultAsync(m => m.SchoolReportId == id);
            if (report == null)
            {
                return NotFound();
            }

            var isAuthorized = await AuthorizationService.AuthorizeAsync(
                                                        User, SchoolReport,
                                                        AuthOperations.Delete);
            if (!isAuthorized.Succeeded)
            {
                return Forbid();
            }
            _context.SchoolReport.Remove(SchoolReport);
                await _context.SaveChangesAsync();

            return RedirectToPage("./Index");
        }
    }
}
