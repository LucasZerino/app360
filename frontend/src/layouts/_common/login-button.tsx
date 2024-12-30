// @mui
import { Theme, SxProps } from '@mui/material/styles';
import Button from '@mui/material/Button';
// routes
import { paths } from '@/routes/paths';
import { RouterLink } from '@/routes/components';
// auth
import { useAuthContext } from '@/auth/hooks';

// ----------------------------------------------------------------------

const loginPaths: Record<string, string> = {
  jwt: paths.auth.login,
};

type Props = {
  sx?: SxProps<Theme>;
};

export default function LoginButton({ sx }: Props) {
  const { method } = useAuthContext();

  const loginPath = loginPaths[method];

  return (
    <Button component={RouterLink} href={loginPath} variant="outlined" sx={{ mr: 1, ...sx }}>
      Login
    </Button>
  );
}
