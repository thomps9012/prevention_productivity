﻿using System.Threading.Tasks;
using prevention_productivity.Models;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Authorization.Infrastructure;

namespace prevention_productivity.Authorization
{
    public class IsEventAdmin
        : AuthorizationHandler<OperationAuthorizationRequirement, Event>
    {
        protected override Task HandleRequirementAsync(
                                    AuthorizationHandlerContext context,
                                   OperationAuthorizationRequirement requirement,
                                   Event resource)
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