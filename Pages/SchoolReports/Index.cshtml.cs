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

        public IList<SchoolReport> SchoolReport { get;set; }
        public IList<ApplicationUser> TeamMember { get; set; }

        public async Task OnGetAsync()
        {
            var reports = from r in _context.SchoolReport
                          select r;
            var teamMembers = from m in _context.Users
                              select m;

            var isAuthorized = User.IsInRole(Constants.AdminRole);

            var currentUserId = UserManager.GetUserId(User);
            
            if(!isAuthorized)
            {
                reports = reports.Where(r => r.TeamMemberId == currentUserId);
            }
            SchoolReport = await reports.ToListAsync();
            TeamMember = await teamMembers.ToListAsync();

        }
    }
}
