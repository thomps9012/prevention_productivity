﻿using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using prevention_productivity.Authorization;
using Microsoft.EntityFrameworkCore;
using prevention_productivity.Data;
using prevention_productivity.Models;

namespace prevention_productivity.Pages.ProductivityLogs
{
    public class EditModel : DI_BasePageModel
    {
        private readonly ApplicationDbContext _context;

        public EditModel(ApplicationDbContext context,
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
            if (id == null)
            {
                return NotFound();
            }

            ProductivityLog = await _context.ProductivityLog.FirstOrDefaultAsync(m => m.LogID == id);

            if (ProductivityLog == null)
            {
                return NotFound();
            }
            var isAuthorized = await AuthorizationService.AuthorizeAsync(
                                                        User, ProductivityLog,
                                                        ProductivityLogOperations.Update);
            if (!isAuthorized.Succeeded)
            {
                return new ForbidResult();
            }
            return Page();
        }

        // To protect from overposting attacks, enable the specific properties you want to bind to.
        // For more details, see https://aka.ms/RazorPagesCRUD.
        public async Task<IActionResult> OnPostAsync()
        {
            if (!ModelState.IsValid)
            {
                return Page();
            }
            var log = await _context.ProductivityLog.FirstOrDefaultAsync(m => m.LogID == ProductivityLog.LogID);

            if (log == null)
            {
                return NotFound();
            }

            var isAuthorized = await AuthorizationService.AuthorizeAsync(
                                                    User, ProductivityLog,
                                                    ProductivityLogOperations.Update);
            if (!isAuthorized.Succeeded)
            {
                ProductivityLog.Status = ApprovalStatus.Pending;
            }


            await _context.SaveChangesAsync();

            return RedirectToPage("./Index");
        }
    }
}