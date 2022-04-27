﻿using System.Threading.Tasks;
using prevention_productivity.Models;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Authorization.Infrastructure;

namespace prevention_productivity.Authorization
{
    public class IsReportAdmin
        : AuthorizationHandler<OperationAuthorizationRequirement, SchoolReport>
    {
        protected override Task HandleRequirementAsync(
                                    AuthorizationHandlerContext context,
                                   OperationAuthorizationRequirement requirement,
                                   SchoolReport resource)
        {
            if (context.User == null)
            {
                return Task.CompletedTask;
            }
            
            if (context.User.IsInRole(Constants.AdminRole))
            {
                context.Succeed(requirement);
            }

            return Task.CompletedTask;
        }
    }
}