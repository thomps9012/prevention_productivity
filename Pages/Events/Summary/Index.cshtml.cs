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

namespace prevention_productivity.Pages.Events.Summary
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

        public IList<EventSummary> EventSummary { get;set; }

        public async Task OnGetAsync()
        {
            var eventSummary = from p in _context.EventSummary.Include(e => e.Event).Include(e => e.TeamMember)
                               select p;
                
                 var currentUserId = UserManager.GetUserId(User);

            var isAuthorized = User.IsInRole(Constants.AdminRole);

            if (!isAuthorized)
            {
                eventSummary = eventSummary.Where(e => e.Status == ApprovalStatus.Approved
                || e.TeamMemberID == currentUserId);
            }
             EventSummary = await eventSummary.ToListAsync();
        }
    }
}
