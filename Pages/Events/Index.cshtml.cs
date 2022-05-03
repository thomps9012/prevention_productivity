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

namespace prevention_productivity.Pages.Events
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

        public IList<Event> Event { get;set; }
        public IList<ApplicationUser> TeamList { get; set; }
        public IList<GrantProgram> Grants { get; set; }
        public IList<EventSummary> EventSummary { get; set; }
        public string DateSort { get; set; }
        public string TeamMemberSearch { get; set; }
        public string GrantSearch { get; set; }
        public string StatusSearch { get; set; }

        public async Task OnGetAsync(string sortOrder,
            string teamMemberSearch,
            string grantSearch,
            string statusSearch)
        {
            DateSort = sortOrder == "asc" ? "date_desc" : "asc";
            var events = from e in _context.Event
                         select e;

            var eventSummaries = from es in _context.EventSummary
                                 select es;

            var teamList = from t in _context.Users
                           select t;
            var grants = from g in _context.GrantProgram
                         select g;

            var currentUserId = UserManager.GetUserId(User);

            var isAuthorized = User.IsInRole(Constants.AdminRole);
            
            if (!isAuthorized)
            {
                events = events.Where(e => e.EventLead == currentUserId);
            }
            if (!string.IsNullOrEmpty(teamMemberSearch))
            {
                events = events.Where(c => c.EventLead == teamMemberSearch);
            }
            if (!string.IsNullOrEmpty(grantSearch))
            {
                events = events.Where(c => c.GrantProgramId == Int64.Parse(grantSearch));
            }
            if (!string.IsNullOrEmpty(statusSearch))
            {
                switch (statusSearch)
                {
                    case "Approved":
                        events = events.Where(c => c.Status == ApprovalStatus.Approved);
                        break;
                    case "Pending":
                        events = events.Where(c => c.Status == ApprovalStatus.Pending);
                        break;
                    case "Rejected":
                        events = events.Where(c => c.Status == ApprovalStatus.Rejected);
                        break;
                }
            }
            switch (sortOrder)
            {
                case "asc":
                    events = events.OrderBy(c => c.EventStart);
                    break;
                case "date_desc":
                    events = events.OrderByDescending(c => c.EventStart);
                    break;
                default:
                    events = events.OrderBy(c => c.EventStart);
                    break;
            }
            Event = await events.ToListAsync();
            TeamList = await teamList.ToListAsync();
            Grants = await grants.ToListAsync();
            EventSummary = await eventSummaries.ToListAsync();
        }
    }
}
