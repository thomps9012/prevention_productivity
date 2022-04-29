#nullable disable
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using prevention_productivity.Authorization;
using Microsoft.EntityFrameworkCore;
using prevention_productivity.Data;
using prevention_productivity.Models;

namespace prevention_productivity.Pages.ProductivityLogs
{
    public class DeleteModel : DI_BasePageModel
    {
        private readonly ApplicationDbContext _context;

        public DeleteModel(ApplicationDbContext context,
            IAuthorizationService authorizationService,
            UserManager<ApplicationUser> userManager)
            : base(context, authorizationService, userManager)
        {
            _context = context;
        }

        [BindProperty]
        public ProductivityLog ProductivityLog { get; set; }
        public ApplicationUser TeamMember { get; set; }

        public async Task<IActionResult> OnGetAsync(int? id)
        {
            ProductivityLog? _log = await _context.ProductivityLog.FirstOrDefaultAsync(m => m.LogID == id);
            if (_log == null)
            {
                return NotFound();
            }

            TeamMember = await _context.Users.FirstOrDefaultAsync(u => u.Id == _log.TeamMemberID);
            ProductivityLog = _log;
            var isAuthorized = User.IsInRole("Admin");
            if (!isAuthorized)
            {
                return Forbid();
            }
            
            return Page();
        }

        public async Task<IActionResult> OnPostAsync(int? id)
        {
            var log = await _context.ProductivityLog.AsNoTracking()
                    .FirstOrDefaultAsync(m => m.LogID == id);
            if (log == null)
            {
                return NotFound();
            }



            var isAuthorized = User.IsInRole("Admin");
            if (!isAuthorized)
            {
                    return Forbid();
                }
                _context.ProductivityLog.Remove(log);
                await _context.SaveChangesAsync();
            

            return RedirectToPage("./Index");
        }
    }
}
