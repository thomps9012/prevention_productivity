using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using prevention_productivity.Data;
using Microsoft.AspNetCore.Authorization;
using prevention_productivity.Authorization;
using prevention_productivity.Models;
using Microsoft.AspNetCore.Identity.UI.Services;
using prevention_productivity.Services;

var builder = WebApplication.CreateBuilder(args);
var services = builder.Services;
var configuration = builder.Configuration;
// grab environment variables


// Add services to the container.
var connectionString = configuration.GetConnectionString("JAWSDB_URL");
var serverVersion = new MySqlServerVersion(new Version(8, 0, 28));
services.AddDbContext<ApplicationDbContext>(options =>
    options.UseMySql(connectionString, serverVersion));
services.AddDatabaseDeveloperPageExceptionFilter();

services.AddDefaultIdentity<ApplicationUser>(
    options => options.SignIn.RequireConfirmedAccount = true)
    .AddRoles<IdentityRole>()
    .AddEntityFrameworkStores<ApplicationDbContext>();

services.AddRazorPages();
services.AddTransient<IEmailSender, EmailSender>();
services.Configure<AuthMessageSenderOptions>(builder.Configuration);

services.AddAuthentication()
    .AddGoogle(options =>
    {
          options.ClientId = Environment.GetEnvironmentVariable("ClientId");
         options.ClientSecret = Environment.GetEnvironmentVariable("ClientSecret");
    });
services.AddAuthorization(options =>
{ 
    options.FallbackPolicy = new AuthorizationPolicyBuilder()
    .RequireAuthenticatedUser()
    .Build();
});

//make sure to register the auth policy handler
services.AddScoped<IAuthorizationHandler, IsTeamMemberHandler>();
services.AddScoped<IAuthorizationHandler, IsEventLead>();
services.AddScoped<IAuthorizationHandler, IsSummaryLead>();
services.AddScoped<IAuthorizationHandler, IsReportHandler>();

services.AddSingleton<IAuthorizationHandler, IsLogAdmin>();
services.AddSingleton<IAuthorizationHandler, IsEventAdmin>();
services.AddSingleton<IAuthorizationHandler, IsSummaryAdmin>();
services.AddSingleton<IAuthorizationHandler, IsReportAdmin>();
services.AddSingleton<IAuthorizationHandler, IsContactAdmin>();

var app = builder.Build();

    app.UseExceptionHandler("/Error");
    // The default HSTS value is 30 days. You may want to change this for production scenarios, see https://aka.ms/aspnetcore-hsts.
    app.UseHsts();


app.UseHttpsRedirection();
app.UseStaticFiles();

app.UseRouting();

app.UseAuthentication();
app.UseAuthorization();

app.MapRazorPages();

app.Run();
