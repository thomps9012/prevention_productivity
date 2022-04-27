using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using prevention_productivity.Data;
using Microsoft.AspNetCore.Authorization;
using prevention_productivity.Authorization;
using prevention_productivity.Models;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
var connectionString = builder.Configuration.GetConnectionString("DefaultConnection");
builder.Services.AddDbContext<ApplicationDbContext>(options =>
    options.UseSqlServer(connectionString));
builder.Services.AddDatabaseDeveloperPageExceptionFilter();

builder.Services.AddDefaultIdentity<ApplicationUser>(
    options => options.SignIn.RequireConfirmedAccount = true)
    .AddRoles<IdentityRole>()
    .AddEntityFrameworkStores<ApplicationDbContext>();

builder.Services.AddRazorPages();

builder.Services.AddAuthorization(options =>
{
    options.FallbackPolicy = new AuthorizationPolicyBuilder()
    .RequireAuthenticatedUser()
    .Build();
});

//make sure to register the auth policy handler
builder.Services.AddScoped<IAuthorizationHandler, IsTeamMemberHandler>();
builder.Services.AddScoped<IAuthorizationHandler, IsEventLead>();
builder.Services.AddScoped<IAuthorizationHandler, IsSummaryLead>();
builder.Services.AddScoped<IAuthorizationHandler, IsReportHandler>();

builder.Services.AddSingleton<IAuthorizationHandler, IsLogAdmin>();
builder.Services.AddSingleton<IAuthorizationHandler, IsEventAdmin>();
builder.Services.AddSingleton<IAuthorizationHandler, IsSummaryAdmin>();
builder.Services.AddSingleton<IAuthorizationHandler, IsReportAdmin>();
builder.Services.AddSingleton<IAuthorizationHandler, IsContactAdmin>();

var app = builder.Build();

using (var scope = app.Services.CreateScope())
{
    var services = scope.ServiceProvider;
    var dbContext = services.GetRequiredService<ApplicationDbContext>();
    dbContext.Database.Migrate();

    var testUserPw = builder.Configuration.GetValue<string>("SeedUserPW");
    await SeedData.Initialize(services, testUserPw);
}
// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseMigrationsEndPoint();
}
else
{
    app.UseExceptionHandler("/Error");
    // The default HSTS value is 30 days. You may want to change this for production scenarios, see https://aka.ms/aspnetcore-hsts.
    app.UseHsts();
}

app.UseHttpsRedirection();
app.UseStaticFiles();

app.UseRouting();

app.UseAuthentication();
app.UseAuthorization();

app.MapRazorPages();

app.Run();
