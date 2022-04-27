using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace prevention_productivity.Data.Migrations
{
    public partial class UpdateGrantModel : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {

            migrationBuilder.AddColumn<string>(
                name: "AwardNumber",
                table: "GrantProgram",
                type: "nvarchar(max)",
                nullable: false,
                defaultValue: "");
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropColumn(
                name: "AwardNumber",
                table: "GrantProgram");

            migrationBuilder.AddColumn<string>(
                name: "Event",
                table: "Contact",
                type: "nvarchar(max)",
                nullable: false,
                defaultValue: "");
        }
    }
}
