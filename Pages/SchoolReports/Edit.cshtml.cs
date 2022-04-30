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

namespace prevention_productivity.Pages.SchoolReports
{
    public class EditModel : DI_BasePageModel
    {
        private readonly ApplicationDbContext _context;

        public EditModel(ApplicationDbContext context,
            IAuthorizationService authorizationService,
            UserManager<ApplicationUser> userManager)
            : base(context, authorizationService, userManager) => _context = context;


        [BindProperty]
        public SchoolReport SchoolReport { get; set; }

        public async Task<IActionResult> OnGetAsync(int? id)
        {
            SchoolReport? report = await _context.SchoolReport.FirstOrDefaultAsync(m => m.SchoolReportId == id);
            if (report == null)
            {
                return NotFound();
            }

            SchoolReport = report;
            if ((await AuthorizationService.AuthorizeAsync(User, report, AuthOperations.Update)).Succeeded)
            {
                return Page();
            } else
            {
                return Forbid();
            }

        }

        // To protect from overposting attacks, enable the specific properties you want to bind to.
        // For more details, see https://aka.ms/RazorPagesCRUD.
        public async Task<IActionResult> OnPostAsync(int id)
        {
           // if (!ModelState.IsValid)
           // {
            //    return Page();
           // }

            var report = await Context.SchoolReport.AsNoTracking().FirstOrDefaultAsync(m => m.SchoolReportId == id);

            if (report == null)
            {
                return NotFound();
            }
            var isAuthorized = await AuthorizationService.AuthorizeAsync(User, report, AuthOperations.Update);
            if (!isAuthorized.Succeeded)
            {
                return Forbid();
            }

                Context.Attach(SchoolReport).State = EntityState.Modified;

                if (SchoolReport.Status == ApprovalStatus.Approved)
                {
                    var canApprove = await AuthorizationService.AuthorizeAsync(User,
                                                                                SchoolReport,
                                                                                AuthOperations.Approve);
                    if (!canApprove.Succeeded)
                    {
                        SchoolReport.Status = ApprovalStatus.Pending;
                    }
                }


                await Context.SaveChangesAsync();

            return RedirectToPage("./Details", new { id = SchoolReport.SchoolReportId });



        }
    }
}