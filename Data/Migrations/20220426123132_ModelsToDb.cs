using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace prevention_productivity.Data.Migrations
{
    public partial class ModelsToDb : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_Comment_ProductivityLog_ProductivityLogLogID",
                table: "Comment");

            migrationBuilder.DropForeignKey(
                name: "FK_Comment_Users_AuthorId",
                table: "Comment");

            migrationBuilder.RenameColumn(
                name: "ProductivityLogLogID",
                table: "Comment",
                newName: "EventId");

            migrationBuilder.RenameColumn(
                name: "CommentId",
                table: "Comment",
                newName: "Id");

            migrationBuilder.RenameIndex(
                name: "IX_Comment_ProductivityLogLogID",
                table: "Comment",
                newName: "IX_Comment_EventId");

            migrationBuilder.AddColumn<int>(
                name: "GrantProgramId",
                table: "Users",
                type: "int",
                nullable: true);

            migrationBuilder.AddColumn<int>(
                name: "GrantProgramId",
                table: "ProductivityLog",
                type: "int",
                nullable: true);

            migrationBuilder.AlterColumn<string>(
                name: "AuthorId",
                table: "Comment",
                type: "nvarchar(450)",
                nullable: true,
                oldClrType: typeof(string),
                oldType: "nvarchar(450)");

            migrationBuilder.CreateTable(
                name: "GrantProgram",
                columns: table => new
                {
                    Id = table.Column<int>(type: "int", nullable: false)
                        .Annotation("SqlServer:Identity", "1, 1"),
                    Name = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Description = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    StartDate = table.Column<DateTime>(type: "datetime2", nullable: false),
                    EndDate = table.Column<DateTime>(type: "datetime2", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_GrantProgram", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "Event",
                columns: table => new
                {
                    Id = table.Column<int>(type: "int", nullable: false)
                        .Annotation("SqlServer:Identity", "1, 1"),
                    Name = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Description = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Location = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    LocationDetail = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    GrantProgramId = table.Column<int>(type: "int", nullable: false),
                    PublicEvent = table.Column<bool>(type: "bit", nullable: false),
                    RSVPRequired = table.Column<bool>(type: "bit", nullable: false),
                    AnnualEvent = table.Column<bool>(type: "bit", nullable: false),
                    NewEvent = table.Column<bool>(type: "bit", nullable: false),
                    EventStart = table.Column<DateTime>(type: "datetime2", nullable: false),
                    Duration = table.Column<double>(type: "float", nullable: false),
                    SetUpTime = table.Column<DateTime>(type: "datetime2", nullable: false),
                    CleanUpTime = table.Column<DateTime>(type: "datetime2", nullable: false),
                    Agenda = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    TargetPopulation = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    AgeGroup = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    PartingGift = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Raffle = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Marketing = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    SpecialOrderGear = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    SepcialOrderDesignElements = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Performance = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Vendors = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    FoodBeverage = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Caterer = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    FoodHeadCount = table.Column<int>(type: "int", nullable: false),
                    OtherSpecialOrder = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    PreventionTeamMembers = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    StaffTimeCommitment = table.Column<double>(type: "float", nullable: false),
                    StaffDuties = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    NORAClients = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Volunteers = table.Column<bool>(type: "bit", nullable: false),
                    VolunteerList = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Supplies = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    EventBudget = table.Column<double>(type: "float", nullable: false),
                    EventAffiliates = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Notes = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    EventLeadId = table.Column<string>(type: "nvarchar(450)", nullable: true),
                    Status = table.Column<int>(type: "int", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Event", x => x.Id);
                    table.ForeignKey(
                        name: "FK_Event_GrantProgram_GrantProgramId",
                        column: x => x.GrantProgramId,
                        principalTable: "GrantProgram",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                    table.ForeignKey(
                        name: "FK_Event_Users_EventLeadId",
                        column: x => x.EventLeadId,
                        principalTable: "Users",
                        principalColumn: "Id");
                });

            migrationBuilder.CreateIndex(
                name: "IX_Users_GrantProgramId",
                table: "Users",
                column: "GrantProgramId");

            migrationBuilder.CreateIndex(
                name: "IX_ProductivityLog_GrantProgramId",
                table: "ProductivityLog",
                column: "GrantProgramId");

            migrationBuilder.CreateIndex(
                name: "IX_Event_EventLeadId",
                table: "Event",
                column: "EventLeadId");

            migrationBuilder.CreateIndex(
                name: "IX_Event_GrantProgramId",
                table: "Event",
                column: "GrantProgramId");

            migrationBuilder.AddForeignKey(
                name: "FK_Comment_Event_EventId",
                table: "Comment",
                column: "EventId",
                principalTable: "Event",
                principalColumn: "Id");

            migrationBuilder.AddForeignKey(
                name: "FK_Comment_Users_AuthorId",
                table: "Comment",
                column: "AuthorId",
                principalTable: "Users",
                principalColumn: "Id");

            migrationBuilder.AddForeignKey(
                name: "FK_ProductivityLog_GrantProgram_GrantProgramId",
                table: "ProductivityLog",
                column: "GrantProgramId",
                principalTable: "GrantProgram",
                principalColumn: "Id");

            migrationBuilder.AddForeignKey(
                name: "FK_Users_GrantProgram_GrantProgramId",
                table: "Users",
                column: "GrantProgramId",
                principalTable: "GrantProgram",
                principalColumn: "Id");
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_Comment_Event_EventId",
                table: "Comment");

            migrationBuilder.DropForeignKey(
                name: "FK_Comment_Users_AuthorId",
                table: "Comment");

            migrationBuilder.DropForeignKey(
                name: "FK_ProductivityLog_GrantProgram_GrantProgramId",
                table: "ProductivityLog");

            migrationBuilder.DropForeignKey(
                name: "FK_Users_GrantProgram_GrantProgramId",
                table: "Users");

            migrationBuilder.DropTable(
                name: "Event");

            migrationBuilder.DropTable(
                name: "GrantProgram");

            migrationBuilder.DropIndex(
                name: "IX_Users_GrantProgramId",
                table: "Users");

            migrationBuilder.DropIndex(
                name: "IX_ProductivityLog_GrantProgramId",
                table: "ProductivityLog");

            migrationBuilder.DropColumn(
                name: "GrantProgramId",
                table: "Users");

            migrationBuilder.DropColumn(
                name: "GrantProgramId",
                table: "ProductivityLog");

            migrationBuilder.RenameColumn(
                name: "EventId",
                table: "Comment",
                newName: "ProductivityLogLogID");

            migrationBuilder.RenameColumn(
                name: "Id",
                table: "Comment",
                newName: "CommentId");

            migrationBuilder.RenameIndex(
                name: "IX_Comment_EventId",
                table: "Comment",
                newName: "IX_Comment_ProductivityLogLogID");

            migrationBuilder.AlterColumn<string>(
                name: "AuthorId",
                table: "Comment",
                type: "nvarchar(450)",
                nullable: false,
                defaultValue: "",
                oldClrType: typeof(string),
                oldType: "nvarchar(450)",
                oldNullable: true);

            migrationBuilder.AddForeignKey(
                name: "FK_Comment_ProductivityLog_ProductivityLogLogID",
                table: "Comment",
                column: "ProductivityLogLogID",
                principalTable: "ProductivityLog",
                principalColumn: "LogID");

            migrationBuilder.AddForeignKey(
                name: "FK_Comment_Users_AuthorId",
                table: "Comment",
                column: "AuthorId",
                principalTable: "Users",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);
        }
    }
}
