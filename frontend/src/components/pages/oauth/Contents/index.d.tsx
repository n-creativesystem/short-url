import type {
  OAuthApplication,
  OAuthApplicationInput,
} from '@/graphql/generated';

export type Data = Omit<OAuthApplication, 'domain' | '__typename'>;

export type Input = Expand<Omit<Data, 'secret'> & OAuthApplicationInput>;
