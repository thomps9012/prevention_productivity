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
        public IList<ApplicationUser> TeamMembers { get; set; }
        public string TeamMemberSearch { get; set; }
        public string StatusSearch { get; set; }

        public async Task OnGetAsync(string teamMemberSearch,
            string statusSearch)
        {
            var eventSummary = from p in _context.EventSummary.Include(e => e.Event).Include(e => e.TeamMember)
                               select p;
                
                 var currentUserId = UserManager.GetUserId(User);
            var teamList = from t in _context.Users
                           select t;

            var isAuthorized = User.IsInRole(Constants.AdminRole);

            if (!isAuthorized)
            {
                eventSummary = eventSummary.Where(e => e.TeamMemberID == currentUserId);
            }
            if (!string.IsNullOrEmpty(teamMemberSearch))
            {
                eventSummary = eventSummary.Where(c => c.TeamMemberID == teamMemberSearch);
            }
            if (!string.IsNullOrEmpty(statusSearch))
            {
                switch (statusSearch)
                {
                    case "Approved":
                        eventSummary = eventSummary.Where(c => c.Status == ApprovalStatus.Approved);
                        break;
                    case "Pending":
                        eventSummary = eventSummary.Where(c => c.Status == ApprovalStatus.Pending);
                        break;
                    case "Rejected":
                        eventSummary = eventSummary.Where(c => c.Status == ApprovalStatus.Rejected);
                        break;
                }
            }
            EventSummary = await eventSummary.ToListAsync();
            TeamMembers = await teamList.ToListAsync();
        }
    }
}
