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

        public async Task OnGetAsync()
        {
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
            Event = await events.ToListAsync();
            TeamList = await teamList.ToListAsync();
            Grants = await grants.ToListAsync();
            EventSummary = await eventSummaries.ToListAsync();
        }
    }
}
