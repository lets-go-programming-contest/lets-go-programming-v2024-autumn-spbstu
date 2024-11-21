package feats

func FindByName(q string) *Feat {
  for _, feat := range Features {
    if feat.Name == q {
      return &feat
    }
  }
  return nil
}
