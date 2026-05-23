import { Outlet } from '@@/exports';
import { useEffect } from 'react';

function BaseLayout() {
  useEffect(() => {}, []);
  return (
    <div>
      <Outlet />
    </div>
  );
}

export default BaseLayout;
