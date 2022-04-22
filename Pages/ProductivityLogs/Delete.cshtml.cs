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
            UserManager<IdentityUser> userManager)
            : base(context, authorizationService, userManager)
        {
            _context = context;
        }

        [BindProperty]
        public ProductivityLog ProductivityLog { get; set; }

        public async Task<IActionResult> OnGetAsync(int? id)
        {
            ProductivityLog = await _context.ProductivityLog.FirstOrDefaultAsync(m => m.LogID == id);
            if (id == null)
            {
                return NotFound();
            }


            if (ProductivityLog == null)
            {
                return NotFound();
            }
            var isAuthorized = await AuthorizationService.AuthorizeAsync(
                                                        User, ProductivityLog,
                                                        ProductivityLogOperations.Delete);
            if (!isAuthorized.Succeeded)
            {
                return new ForbidResult();
            }
            
            return Page();
        }

        public async Task<IActionResult> OnPostAsync(int? id)
        {
            if (id == null)
            {
                return NotFound();
            }

            ProductivityLog = await _context.ProductivityLog.FindAsync(id);

            if (ProductivityLog != null)
            {
                var isAuthorized = await AuthorizationService.AuthorizeAsync(
                                                        User, ProductivityLog,
                                                        ProductivityLogOperations.Delete);
                if(!isAuthorized.Succeeded)
                {
                    return new ForbidResult();
                }
                _context.ProductivityLog.Remove(ProductivityLog);
                await _context.SaveChangesAsync();
            }

            return RedirectToPage("./Index");
        }
    }
}
