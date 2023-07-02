import { Url } from '@/components/pages/shorts/graphql';

type Data = {
  url: string;
} & Omit<Url, 'url' | '__typename'>;

export type Input = Expand<Pick<Data, 'url'>>;
