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
    public class DetailsModel : DI_BasePageModel
    {
        private readonly ApplicationDbContext _context;

        public DetailsModel(ApplicationDbContext context,
            IAuthorizationService authorizationService,
            UserManager<ApplicationUser> userManager)
            : base(context, authorizationService, userManager)
        {
            _context = context;
        }

        public SchoolReport SchoolReport { get; set; }

        public async Task<IActionResult> OnGetAsync(int? id)
        {
            SchoolReport? _report = await _context.SchoolReport
                .FirstOrDefaultAsync(m => m.SchoolReportId == id);
            if (_report == null)
            {
                return NotFound();
            }

            SchoolReport = _report;

            var isAdmin = User.IsInRole(Constants.AdminRole);

            var currentUserId = UserManager.GetUserId(User);

            if (!isAdmin
                && currentUserId != SchoolReport.TeamMemberId
                && SchoolReport.Status != ApprovalStatus.Approved)
            {
                return Forbid();
            }
            return Page();
        }
        public async Task<IActionResult> OnPostAsync(int id, ApprovalStatus status)
        {
            var report = await _context.SchoolReport.FirstOrDefaultAsync(m => m.SchoolReportId == id);

            if (report == null)
            {
                return NotFound();
            }

            var operation = (status == ApprovalStatus.Approved)
                ? AuthOperations.Approve
                : AuthOperations.Reject;

            var isAuthorized = await AuthorizationService.AuthorizeAsync(User, report, operation);

            if (!isAuthorized.Succeeded)
            {
                return Forbid();
            }
            report.Status = status;
            _context.SchoolReport.Update(report);
            await _context.SaveChangesAsync();
            return RedirectToPage("./Index");
        }
    }
}