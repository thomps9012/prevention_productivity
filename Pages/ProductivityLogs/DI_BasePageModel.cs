using prevention_productivity.Data;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Microsoft.AspNetCore.Identity;
namespace prevention_productivity.Pages.ProductivityLogs
{
    public class DI_BasePageModel : PageModel
    {
        protected ApplicationDbContext Context;
        protected IAuthorizationService AuthorizationService;
        protected UserManager<IdentityUser> UserManager;
        public DI_BasePageModel(
            ApplicationDbContext context,
            IAuthorizationService authorizationService,
            UserManager<IdentityUser> userManager)
        {
            Context = context;
            AuthorizationService = authorizationService;
            UserManager = userManager;
        }
    }
}
