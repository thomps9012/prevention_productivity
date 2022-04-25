using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using prevention_productivity.Data;
using prevention_productivity.Models;
using prevention_productivity.Authorization;

namespace prevention_productivity.Pages.ProductivityLogs
{
    public class CreateModel : DI_BasePageModel
    {
        private readonly ApplicationDbContext _context;

        public CreateModel(ApplicationDbContext context,
            IAuthorizationService authorizationService,
            UserManager<IdentityUser> userManager)
            : base(context, authorizationService, userManager)
        {
            _context = context;
        }

        public IActionResult OnGet()
        {
            return Page();
        }

        [BindProperty]
        public ProductivityLog ProductivityLog { get; set; }

        // To protect from overposting attacks, see https://aka.ms/RazorPagesCRUD
        public async Task<IActionResult> OnPostAsync()
        {
           
            ProductivityLog.TeamMemberID = UserManager.GetUserId(User);
            ProductivityLog.Status = ApprovalStatus.Pending;

            var isAuthorized = await AuthorizationService.AuthorizeAsync(
                                                        User, ProductivityLog,
                                                        ProductivityLogOperations.Create);
            if (!isAuthorized.Succeeded)
            {
                return new ForbidResult();
            }
            _context.ProductivityLog.Add(ProductivityLog);
            await _context.SaveChangesAsync();

            return RedirectToPage("./Index");
        }
    }
}
