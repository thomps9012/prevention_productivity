using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace prevention_productivity.Data.Migrations
{
    public partial class UpdatedReportModel : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_SchoolReport_Users_TeamMemberId",
                table: "SchoolReport");

            migrationBuilder.DropIndex(
                name: "IX_SchoolReport_TeamMemberId",
                table: "SchoolReport");

            migrationBuilder.AlterColumn<string>(
                name: "TeamMemberId",
                table: "SchoolReport",
                type: "nvarchar(max)",
                nullable: false,
                defaultValue: "",
                oldClrType: typeof(string),
                oldType: "nvarchar(450)",
                oldNullable: true);
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.AlterColumn<string>(
                name: "TeamMemberId",
                table: "SchoolReport",
                type: "nvarchar(450)",
                nullable: true,
                oldClrType: typeof(string),
                oldType: "nvarchar(max)");

            migrationBuilder.CreateIndex(
                name: "IX_SchoolReport_TeamMemberId",
                table: "SchoolReport",
                column: "TeamMemberId");

            migrationBuilder.AddForeignKey(
                name: "FK_SchoolReport_Users_TeamMemberId",
                table: "SchoolReport",
                column: "TeamMemberId",
                principalTable: "Users",
                principalColumn: "Id");
        }
    }
}
