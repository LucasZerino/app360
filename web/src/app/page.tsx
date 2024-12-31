// sections
// ----------------------------------------------------------------------

import { Box, Container, Typography, Button, Stack } from '@mui/material';
import Link from 'next/link';

export const metadata = {
  title: 'Bolt 360 App - O app completo para o seu negócio',
  description: 'Gerencie seu negócio de forma simples e eficiente com o Bolt 360',
};

export default function HomePage() {
  return (
    <Container>
      <Box
        sx={{
          py: 12,
          textAlign: 'center',
          minHeight: '100vh',
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'center',
        }}
      >
        <Typography variant="h1" sx={{ mb: 3 }}>
          Bem-vindo ao Bolt 360
        </Typography>

        <Typography variant="h4" sx={{ color: 'text.secondary', mb: 5 }}>
          A solução completa para o seu negócio
        </Typography>

        <Typography sx={{ color: 'text.secondary', mb: 5 }}>
          Gerencie vendas, estoque, clientes e muito mais em um só lugar.
          Simplifique sua gestão e aumente seus resultados.
        </Typography>

        <Stack direction="row" spacing={2} justifyContent="center">
          <Link href="/dashboard" style={{ textDecoration: 'none' }}>
            <Button size="large" variant="contained">
              Acessar Dashboard
            </Button>
          </Link>
          
          <Link href="/auth/login" style={{ textDecoration: 'none' }}>
            <Button size="large" variant="outlined">
              Login
            </Button>
          </Link>
        </Stack>
      </Box>
    </Container>
  );
}
