import { useEffect, useRef } from 'react';
import {
  Location,
  useLocation as useBaseLocation,
  useParams as useBaseParams,
  useNavigate,
} from 'react-router-dom';
export * from './index.tsx';

export const useRouter = () => {
  const navigation = useNavigate();
  return {
    push: navigation,
  };
};

export const useParams = () => {
  return useBaseParams();
};

export const useLocation = () => {
  return useBaseLocation();
};

export const usePathname = () => {
  const location = useLocation();
  return location.pathname;
};

export const useLocationChange = (callback: (location: Location) => void) => {
  const refCallback = useRef<undefined | ((location: Location) => void)>();
  const location = useLocation();
  useEffect(() => {
    refCallback.current = callback;
  }, [callback]);
  useEffect(() => {
    if (refCallback.current) {
      refCallback.current(location);
    }
  }, [location]);
};
