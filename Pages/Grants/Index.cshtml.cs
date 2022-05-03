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
using prevention_productivity.Data;
using prevention_productivity.Models;
using prevention_productivity.Pages.ProductivityLogs;

namespace prevention_productivity.Pages.Grants
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

        public IList<GrantProgram> GrantProgram { get;set; }
        public string StartDateSort { get; set; }
        public string EndDateSort { get; set; }
        public async Task OnGetAsync(string endDateSort, string startDateSort)
        {
            EndDateSort = endDateSort == "asc" ? "date_desc" : "asc";
            StartDateSort = startDateSort == "asc" ? "date_desc" : "asc";

            var grants = from g in _context.GrantProgram
                         select g;

            switch (endDateSort)
            {
                case "asc":
                    grants = grants.OrderBy(c => c.EndDate);
                    break;
                case "date_desc":
                    grants = grants.OrderByDescending(c => c.EndDate);
                    break;
                default:
                    grants = grants.OrderBy(c => c.EndDate);
                    break;
            }
            switch (startDateSort)
            {
                case "asc":
                    grants = grants.OrderBy(c => c.StartDate);
                    break;
                case "date_desc":
                    grants = grants.OrderByDescending(c => c.StartDate);
                    break;
                default:
                    grants = grants.OrderBy(c => c.StartDate);
                    break;
            }
            GrantProgram = await grants.ToListAsync();
        }
    }
}
