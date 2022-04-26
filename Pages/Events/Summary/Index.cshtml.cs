#nullable disable
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Microsoft.EntityFrameworkCore;
using prevention_productivity.Data;
using prevention_productivity.Models;

namespace prevention_productivity.Pages.Events.Summary
{
    public class IndexModel : PageModel
    {
        private readonly prevention_productivity.Data.ApplicationDbContext _context;

        public IndexModel(prevention_productivity.Data.ApplicationDbContext context)
        {
            _context = context;
        }

        public IList<EventSummary> EventSummary { get;set; }

        public async Task OnGetAsync()
        {
            EventSummary = await _context.EventSummary
                .Include(e => e.Event)
                .Include(e => e.TeamMember).ToListAsync();
        }
    }
}
