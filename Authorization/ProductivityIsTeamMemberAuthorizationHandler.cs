using prevention_productivity.Models;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Authorization.Infrastructure;
using Microsoft.AspNetCore.Identity;
using System.Threading.Tasks;

namespace prevention_productivity.Authorization
{
    public class IsTeamMemberHandler
        : AuthorizationHandler<OperationAuthorizationRequirement, ProductivityLog>
    {

            UserManager<ApplicationUser> _userManager;
        public IsTeamMemberHandler(UserManager<ApplicationUser> userManager)
        {
            _userManager = userManager;
        }
        
        protected override Task
            HandleRequirementAsync(AuthorizationHandlerContext context,
                                   OperationAuthorizationRequirement requirement,
                                   ProductivityLog resource)
        {
            if (context.User == null || resource == null)
            {
                return Task.CompletedTask;
            }
            
            if (requirement.Name != Constants.CreateOperationName &&
                requirement.Name != Constants.ReadOperationName &&
                requirement.Name != Constants.UpdateOperationName &&
                requirement.Name != Constants.DeleteOperationName)
            {
                return Task.CompletedTask;
            }

            if (resource.TeamMemberID == _userManager.GetUserId(context.User))
            {
                context.Succeed(requirement);
            }

            return Task.CompletedTask;
        }
    }
}