using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace prevention_productivity.Data.Migrations
{
    public partial class FullNameColumn : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_Event_Users_EventLeadId",
                table: "Event");

            migrationBuilder.DropIndex(
                name: "IX_Event_EventLeadId",
                table: "Event");

            migrationBuilder.DropColumn(
                name: "EventLeadId",
                table: "Event");

            migrationBuilder.AddColumn<string>(
                name: "EventLead",
                table: "Event",
                type: "nvarchar(max)",
                nullable: false,
                defaultValue: "");
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropColumn(
                name: "EventLead",
                table: "Event");

            migrationBuilder.AddColumn<string>(
                name: "EventLeadId",
                table: "Event",
                type: "nvarchar(450)",
                nullable: true);

            migrationBuilder.CreateIndex(
                name: "IX_Event_EventLeadId",
                table: "Event",
                column: "EventLeadId");

            migrationBuilder.AddForeignKey(
                name: "FK_Event_Users_EventLeadId",
                table: "Event",
                column: "EventLeadId",
                principalTable: "Users",
                principalColumn: "Id");
        }
    }
}
