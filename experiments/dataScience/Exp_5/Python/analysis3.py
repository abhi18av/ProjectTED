import textacy


textacy.preprocess_text(doc.text, lowercase=True, no_punct=True)[:70]
textacy.text_utils.keyword_in_context(doc.text, 'America', window_width=35)


list(textacy.extract.ngrams(
    doc, 2, filter_stops=True, filter_punct=True, filter_nums=False))[:15]