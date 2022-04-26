﻿using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Identity.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore;
using prevention_productivity.Models;

namespace prevention_productivity.Data
{
    public class ApplicationDbContext : 
        IdentityDbContext<ApplicationUser>
    {
        public ApplicationDbContext(DbContextOptions<ApplicationDbContext> options)
            : base(options)
        {
        }
        public DbSet<GrantProgram> GrantProgram { get; set; }
        public DbSet<ProductivityLog> ProductivityLog { get; set; }
        public DbSet<Event> Event { get; set; }
        public DbSet<EventSummary> EventSummary { get; set; }
        public DbSet<Contact> Conttact { get; set; }
        public DbSet<SchoolReport> SchoolReport { get; set; }
        public DbSet<Comment> Comment { get; set; }

        protected override void OnModelCreating(ModelBuilder builder)
        {
            base.OnModelCreating(builder);
            // Customize the ASP.NET Identity model and override the defaults if needed.
            // For example, you can rename the ASP.NET Identity table names and more.
            // Add your customizations after calling base.OnModelCreating(builder);
            builder.Entity<ApplicationUser>(b =>
            {
                b.ToTable("Users");
            });
        }

    }
}