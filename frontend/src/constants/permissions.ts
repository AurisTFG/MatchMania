import { Permission } from 'types/enums/permission';

export const AllPermissions: Permission[] = [
  Permission.Admin,

  Permission.ViewLeague,
  Permission.ManageLeague,

  Permission.ViewTeam,
  Permission.ManageTeam,

  Permission.ViewResult,
  Permission.ManageResult,

  Permission.ViewUser,
  Permission.ManageUser,

  Permission.ViewRole,
  Permission.ManageRole,

  Permission.ViewQueue,
  Permission.ManageQueue,

  Permission.ViewMatch,
  Permission.ManageMatch,
];
