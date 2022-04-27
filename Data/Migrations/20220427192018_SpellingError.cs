using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace prevention_productivity.Data.Migrations
{
    public partial class SpellingError : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.RenameColumn(
                name: "SepcialOrderDesignElements",
                table: "Event",
                newName: "SpecialOrderDesignElements");
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.RenameColumn(
                name: "SpecialOrderDesignElements",
                table: "Event",
                newName: "SepcialOrderDesignElements");
        }
    }
}
