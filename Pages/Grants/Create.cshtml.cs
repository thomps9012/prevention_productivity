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
using prevention_productivity.Authorization;
using prevention_productivity.Data;
using prevention_productivity.Models;
using prevention_productivity.Pages.ProductivityLogs;

namespace prevention_productivity.Pages.Grants
{
    public class CreateModel : DI_BasePageModel
    {
        private readonly ApplicationDbContext _context;

        public CreateModel(ApplicationDbContext context,
            IAuthorizationService authorizationService,
            UserManager<ApplicationUser> userManager)
            : base(context, authorizationService, userManager)
        {
            _context = context;
        }

        public IActionResult OnGet()
        {
            var isAuthorized = User.IsInRole(Constants.AdminRole);
            if (!isAuthorized)
            {
                return Forbid();
            }
            return Page();
        }

        [BindProperty]
        public GrantProgram GrantProgram { get; set; }

        // To protect from overposting attacks, see https://aka.ms/RazorPagesCRUD
        public async Task<IActionResult> OnPostAsync()
        {
            var isAuthorized = User.IsInRole(Constants.AdminRole);
            if (!isAuthorized)
            {
                return Forbid();
            }            

            _context.GrantProgram.Add(GrantProgram);
            await _context.SaveChangesAsync();

            return RedirectToPage("./Index");
        }
    }
}
