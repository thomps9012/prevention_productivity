# FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build
# WORKDIR /prevention_productivity

# COPY *.sln .
# COPY prevention_productivity/*.csproj ./prevention_productivity/
# RUN dotnet restore

# COPY prevention_productivity/. ./prevention_productivity/
# WORKDIR /prevention_productivity
# RUN dotnet public -c release -o /prevention_productivity/bin/Release --no-restore

# FROM mcr.microsoft.com/dotnet/aspnet:6.0
# WORKDIR /prevention_productivity
# COPY --from=build /prevention_productivity ./
# ENTRYPOINT [ "dotnet", "prevention_productivity.dll" ]

FROM mcr.microsoft.com/dotnet/aspnet:6.0 AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443

FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build
WORKDIR /src
COPY ["prevention_productivity.csproj", "."]
RUN dotnet restore "./prevention_productivity.csproj"
COPY . .
WORKDIR "/src/."
RUN dotnet build "prevention_productivity.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "prevention_productivity.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "prevention_productivity.dll"]