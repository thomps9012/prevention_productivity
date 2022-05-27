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

var connectionString = "server=localhost;user=root;password=root;database=prevention_productivity;";

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

services.ConfigureApplicationCookie(o => {
    o.ExpireTimeSpan = TimeSpan.FromDays(5);
    o.SlidingExpiration = true;
});

services.Configure<DataProtectionTokenProviderOptions>(o =>
    o.TokenLifespan = TimeSpan.FromHours(3));

services.AddAuthentication()
    .AddGoogle(googleOptions =>
    {
          googleOptions.ClientId = configuration["Authentication:Google:ClientId"];
         googleOptions.ClientSecret = configuration["Authentication:Google:ClientSecret"];
    });

services.AddAuthorization(options =>
{ 
    options.FallbackPolicy = new AuthorizationPolicyBuilder()
    .RequireAuthenticatedUser()
    .Build();
});

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
app.UseHsts();


app.UseHttpsRedirection();
app.UseStaticFiles();

app.UseRouting();

app.UseAuthentication();
app.UseAuthorization();

app.MapRazorPages();

app.Run();
