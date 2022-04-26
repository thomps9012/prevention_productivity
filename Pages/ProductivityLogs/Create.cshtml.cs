using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using prevention_productivity.Data;
using prevention_productivity.Models;
using prevention_productivity.Authorization;
using Microsoft.EntityFrameworkCore;

namespace prevention_productivity.Pages.ProductivityLogs
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
            return Page();
        }

        [BindProperty]
        public ProductivityLog ProductivityLog { get; set; }
        public IList<GrantProgram> GrantPrograms { get; set; }

        // To protect from overposting attacks, see https://aka.ms/RazorPagesCRUD
        public async Task OnGetAsync()
        {
            GrantPrograms = await _context.GrantProgram.ToListAsync();
        }
        public async Task<IActionResult> OnPostAsync()
        {
           
            ProductivityLog.TeamMemberID = UserManager.GetUserId(User);
            ProductivityLog.Status = ApprovalStatus.Pending;

            var isAuthorized = await AuthorizationService.AuthorizeAsync(
                                                        User, ProductivityLog,
                                                        AuthOperations.Create);
            if (!isAuthorized.Succeeded)
            {
                return Forbid();
            }
            _context.ProductivityLog.Add(ProductivityLog);
            await _context.SaveChangesAsync();

            return RedirectToPage("./Index");
        }
    }
}
