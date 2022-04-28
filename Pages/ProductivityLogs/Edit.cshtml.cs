using Microsoft.AspNetCore.Authorization;
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
            UserManager<ApplicationUser> userManager)
            : base(context, authorizationService, userManager) => _context = context;

        [BindProperty]
        public ProductivityLog ProductivityLog { get; set; }
        public IList<GrantProgram> GrantPrograms { get; set; }

        public async Task<IActionResult> OnGetAsync(int id)
        {
           
            ProductivityLog? log = await _context.ProductivityLog
                .FirstOrDefaultAsync(m => m.LogID == id);
            if (log == null)
            {
                return NotFound();
            }

            ProductivityLog = log;
            GrantPrograms = await _context.GrantProgram.ToListAsync();

            if ((await AuthorizationService.AuthorizeAsync(User, log, AuthOperations.Update)).Succeeded)
            {
                return Page();
            } else
            {   
                return Forbid();
            }
        }

        // To protect from overposting attacks, enable the specific properties you want to bind to.
        // For more details, see https://aka.ms/RazorPagesCRUD.
        public async Task<IActionResult> OnPostAsync(int id)
        {
            

            var log = await Context.ProductivityLog.AsNoTracking().FirstOrDefaultAsync(m => m.LogID == id);

            if (log == null)
            {
                return NotFound();
            }
            var isAuthorized = await AuthorizationService.AuthorizeAsync(User, log, AuthOperations.Update);
            if (!isAuthorized.Succeeded)
            {
                return Forbid();
            }
            
            Context.Attach(ProductivityLog).State = EntityState.Modified;

                if (ProductivityLog.Status == ApprovalStatus.Approved)
                {
                    var canApprove = await AuthorizationService.AuthorizeAsync(User,
                                                                                ProductivityLog,
                                                                                AuthOperations.Approve);
                    if (!canApprove.Succeeded)
                    {
                        ProductivityLog.Status = ApprovalStatus.Pending;
                    }
                }


                await Context.SaveChangesAsync();

                return RedirectToPage("./Index");

            
        }
    }
}
