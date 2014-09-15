package easter

import (
  "errors"
  "github.com/goodsign/monday"
)

const error_mess string = "Locale not supported."

func getLocale(locale string) (monday.Locale, error) {
  switch locale {
    case "sv_SE":
      return monday.LocaleSvSE, nil
    case "en_US":
      return monday.LocaleEnUS, nil
    case "da_DK":
      return monday.LocaleDaDK, nil
    case "de_DE":
      return monday.LocaleDeDE, nil
    case "es_ES":
      return monday.LocaleEsES, nil
    case "fi_FI":
      return monday.LocaleFiFI, nil
    case "fr_FR":
      return monday.LocaleFrFR, nil
    case "hu_HU":
      return monday.LocaleHuHU, nil
    case "it_IT":
      return monday.LocaleItIT, nil
    case "nb_NO":
      return monday.LocaleNbNO, nil
    case "nl_BE":
      return monday.LocaleNlBE, nil
    case "nn_NO":
      return monday.LocaleNnNO, nil
    case "pt_BR":
      return monday.LocalePtBR, nil
    case "pt_PT":
      return monday.LocalePtPT, nil
    case "ro_RO":
      return monday.LocaleRoRO, nil
    case "ru_RU":
      return monday.LocaleRuRU, nil
    case "tr_TR":
      return monday.LocaleTrTR, nil
    case "zh_ZH":
      return monday.LocaleTrTR, nil
    default:
      return monday.LocaleEnUS, errors.New(error_mess)
  }
}
