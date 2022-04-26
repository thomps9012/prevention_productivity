using prevention_productivity.Models;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Authorization.Infrastructure;
using Microsoft.AspNetCore.Identity;
using System.Threading.Tasks;

namespace prevention_productivity.Authorization
{
    public class IsSummaryLead
        : AuthorizationHandler<OperationAuthorizationRequirement, EventSummary>
    {

            UserManager<ApplicationUser> _userManager;
        public IsSummaryLead(UserManager<ApplicationUser> userManager)
        {
            _userManager = userManager;
        }
        
        protected override Task
            HandleRequirementAsync(AuthorizationHandlerContext context,
                                   OperationAuthorizationRequirement requirement,
                                   EventSummary resource)
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

            if (resource.EventLead == _userManager.GetUserId(context.User))
            {
                context.Succeed(requirement);
            }

            return Task.CompletedTask;
        }
    }
}